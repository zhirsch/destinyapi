package destinyapi

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey  string
	authURL *url.URL

	httpClient *http.Client

	authToken    token
	refreshToken token
}

func NewClient(apiKey, authURL string) (*Client, error) {
	au, err := url.ParseRequestURI(authURL)
	if err != nil {
		return nil, err
	}
	return &Client{
		apiKey:     apiKey,
		authURL:    au,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) get(req Request, resp Response) error {
	httpReq, err := http.NewRequest("GET", req.URL(), nil)
	if err != nil {
		return err
	}
	httpReq.Header.Add("X-API-Key", c.apiKey)
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.authToken.value))

	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response: %v", httpResp.StatusCode)
	}

	if err := decode(resp, httpResp.Body); err != nil {
		return err
	}
	if resp.GetHeader().ErrorCode != 1 {
		return fmt.Errorf("bad message: %v", resp)
	}
	return nil
}

func (c *Client) GetBungieNetUser() (*GetBungieNetUserResponse, error) {
	req := new(GetBungieNetUserRequest)
	resp := new(GetBungieNetUserResponse)
	if err := c.get(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) GetBungieAccount(membershipID string) (*GetBungieAccountResponse, error) {
	req := &GetBungieAccountRequest{membershipID}
	resp := new(GetBungieAccountResponse)
	if err := c.get(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Client) MyCharacterVendorData(characterHash, vendorHash string) (*MyCharacterVendorDataResponse, error) {
	req := &MyCharacterVendorDataRequest{characterHash, vendorHash}
	resp := new(MyCharacterVendorDataResponse)
	if err := c.get(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
