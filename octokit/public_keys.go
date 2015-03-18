package octokit

import (
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

var (
	CurrentPublicKeyUrl = Hyperlink("/user/keys{/id}")
	PublicKeyUrl        = Hyperlink("/users/{user}/keys")
)

// Create a PublicKeysService
func (c *Client) PublicKeys() (k *PublicKeysService) {
	k = &PublicKeysService{client: c}
	return
}

// A service to return user public keys
type PublicKeysService struct {
	client *Client
}

// Get a list of keys for the user
func (k *PublicKeysService) All(uri *Hyperlink, params M) (keys []Key, result *Result) {
	if uri == nil {
		uri = &CurrentPublicKeyUrl // Default url
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = k.client.get(url, &keys)
	return
}

// Get a the data for one key for the current user
func (k *PublicKeysService) One(uri *Hyperlink, params M) (key *Key, result *Result) {
	if uri == nil {
		uri = &CurrentPublicKeyUrl // Default url
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = k.client.get(url, &key)
	return
}

// Creates a new public key for the current user
func (k *PublicKeysService) Create(uri *Hyperlink, uriParams M, inputParams interface{}) (key *Key, result *Result) {
	if uri == nil {
		uri = &CurrentPublicKeyUrl
	}

	url, err := uri.Expand(uriParams)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = k.client.post(url, inputParams, &key)
	return
}

// Removes a public key for the current user
func (k *PublicKeysService) Delete(uri *Hyperlink, params M) (success bool, result *Result) {
	if uri == nil {
		uri = &CurrentPublicKeyUrl
	}

	url, err := uri.Expand(params)
	if err != nil {
		return false, &Result{Err: err}
	}

	result = k.client.delete(url, nil, nil)
	success = (result.Response.StatusCode == 204)
	return
}

type Key struct {
	*hypermedia.HALResource

	Id        int        `json:"id,omitempty"`
	Key       string     `json:"key,omitempty"`
	URL       string     `json:"url,omitempty"`
	Title     string     `json:"title,omitempty"`
	Verified  bool       `json:"verified,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}