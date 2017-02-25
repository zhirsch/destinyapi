package destinyapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Encode marshals a request to bytes for sending to an API endpoint.
func encode(i interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	return buf, err
}

// Decode unmarshals a response from an API endpoint.
func decode(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

type Request interface {
	URL() string
}

type Response interface {
	GetHeader() *Header
}

// A Header contains comment fields in a response from the Bungie API.
type Header struct {
	ErrorCode       int         `json:"ErrorCode"`
	ErrorStatus     string      `json:"ErrorStatus"`
	Message         string      `json:"Message"`
	MessageData     interface{} `json:"MessageData"`
	ThrottleSeconds int         `json:"ThrottleSeconds"`
}

type GetBungieNetUserRequest struct{}

func (*GetBungieNetUserRequest) URL() string {
	return "https://www.bungie.net/Platform/User/GetBungieNetUser/"
}

type GetBungieNetUserResponse struct {
	Header
	Response struct {
		User struct {
			MembershipID string `json:"membershipId"`
			DisplayName  string `json:"displayName"`
		} `json:"user"`
	} `json:"Response"`
}

func (r *GetBungieNetUserResponse) GetHeader() *Header {
	return &r.Header
}

type GetBungieAccountRequest struct {
	MembershipID string
}

func (r *GetBungieAccountRequest) URL() string {
	return fmt.Sprintf("https://www.bungie.net/Platform/User/GetBungieAccount/%s/2/", r.MembershipID)
}

type GetBungieAccountResponse struct {
	Header
	Response struct {
		DestinyAccounts []struct {
			Characters []struct {
				CharacterID    string `json:"characterId"`
				CharacterClass struct {
					ClassName string `json:"className"`
				} `json:"characterClass"`
			} `json:"characters"`
		} `json:"destinyAccounts"`
	} `json:"Response"`
}

func (r *GetBungieAccountResponse) GetHeader() *Header {
	return &r.Header
}

// A GetAccessTokensFromCodeRequest is a request for the GetAccessTokensFromCode
// method.
type GetAccessTokensFromCodeRequest struct {
	Code string `json:"code"`
}

// URL returns the endpoint for the GetAccessTokensFromCode method.
func (GetAccessTokensFromCodeRequest) URL() string {
	return "https://www.bungie.net/Platform/App/GetAccessTokensFromCode/"
}

// A GetAccessTokensFromCodeResponse is a response from the
// GetAccessTokensFromCode method.
type GetAccessTokensFromCodeResponse struct {
	Header
	Response struct {
		AccessToken struct {
			Value   string `json:"value"`
			ReadyIn int    `json:"readyIn"`
			Expires int    `json:"expires"`
		} `json:"accessToken"`
		RefreshToken struct {
			Value   string `json:"value"`
			ReadyIn int    `json:"readyIn"`
			Expires int    `json:"expires"`
		} `json:"refreshToken"`
	}
	Scope int `json:"scope"`
}

func (r *GetAccessTokensFromCodeResponse) GetHeader() *Header {
	return &r.Header
}

type MyCharacterVendorDataRequest struct {
	CharacterHash string
	VendorHash    string
}

func (r *MyCharacterVendorDataRequest) URL() string {
	return fmt.Sprintf("https://www.bungie.net/Platform/Destiny/2/MyAccount/Character/%v/Vendor/%v/?definitions=true", r.CharacterHash, r.VendorHash)
}

type MyCharacterVendorDataResponse struct {
	Header
	Response struct {
		Data struct {
			VendorHash         int `json:"vendorHash"`
			SaleItemCategories []struct {
				CategoryTitle string `json:"categoryTitle"`
				SaleItems     []struct {
					FailureIndexes []int `json:"failureIndexes"`
					Item           struct {
						ItemHash int `json:"itemHash"`
					} `json:"item"`
				} `json:"saleItems"`
			} `json:"saleItemCategories"`
		} `json:"data"`
		Definitions struct {
			Items map[string]struct {
				ItemHash      int    `json:"itemHash"`
				ItemName      string `json:"itemName"`
				Icon          string `json:"icon"`
				SecondaryIcon string `json:"secondaryIcon"`
			} `json:"items"`
			VendorDetails map[string]struct {
				FailureStrings []string `json:"failureStrings"`
			} `json:"vendorDetails"`
		} `json:"definitions"`
	} `json:"Response"`
}

func (r *MyCharacterVendorDataResponse) GetHeader() *Header {
	return &r.Header
}
