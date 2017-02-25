package destinyapi

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type token struct {
	value   string
	ready   time.Time
	expires time.Time
}

func (t token) isReady() bool {
	return time.Now().After(t.ready)
}

func (t token) isExpired() bool {
	return time.Now().After(t.expires)
}

func (c *Client) Authenticate(w http.ResponseWriter, r *http.Request) bool {
	if !c.authToken.isExpired() {
		return true
	}

	u := r.URL
	u.Scheme = "https"
	u.Host = r.Host

	var au url.URL = *c.authURL
	q := au.Query()
	q.Set("state", u.String())
	au.RawQuery = q.Encode()

	http.Redirect(w, r, au.String(), http.StatusSeeOther)
	return false
}

func (c *Client) HandleBungieAuthCallback(w http.ResponseWriter, r *http.Request) {
	// Validate the incoming query.
	query := r.URL.Query()
	if _, ok := query["code"]; !ok {
		http.Error(w, fmt.Sprintf("no 'code' in request: %s", r.URL),
			http.StatusInternalServerError)
		return
	} else if len(query["code"]) != 1 {
		http.Error(w, fmt.Sprintf("multiple 'code' in request: %s", r.URL),
			http.StatusInternalServerError)
		return
	} else if _, ok := query["state"]; !ok {
		http.Error(w, fmt.Sprintf("no `state` in request: %s", r.URL),
			http.StatusInternalServerError)
		return
	} else if len(query["state"]) != 1 {
		http.Error(w, fmt.Sprintf("too many `state` in request: %s", r.URL),
			http.StatusInternalServerError)
		return
	}

	// Prepare the request body.
	req := GetAccessTokensFromCodeRequest{Code: query["code"][0]}
	body, err := encode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the HTTP request.
	httpReq, err := http.NewRequest("POST", req.URL(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpReq.Header.Add("Content-Type", "application/json")
	httpReq.Header.Add("X-API-Key", c.apiKey)
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if httpResp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("bad response: %v", httpResp.StatusCode),
			http.StatusInternalServerError)
		return
	}

	// Parse the HTTP response.
	var resp GetAccessTokensFromCodeResponse
	if err := decode(&resp, httpResp.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if resp.ErrorCode != 1 {
		http.Error(w, fmt.Sprintf("bad message: %+v", resp), http.StatusInternalServerError)
		return
	}

	// Create the tokens.
	now := time.Now()
	c.authToken = token{
		value:   resp.Response.AccessToken.Value,
		ready:   now.Add(time.Duration(resp.Response.AccessToken.ReadyIn) * time.Second),
		expires: now.Add(time.Duration(resp.Response.AccessToken.Expires) * time.Second),
	}
	c.refreshToken = token{
		value:   resp.Response.RefreshToken.Value,
		ready:   now.Add(time.Duration(resp.Response.RefreshToken.ReadyIn) * time.Second),
		expires: now.Add(time.Duration(resp.Response.RefreshToken.Expires) * time.Second),
	}

	// Redirect to the original URL.
	http.Redirect(w, r, query["state"][0], http.StatusSeeOther)
}
