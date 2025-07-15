package tokensvc

type CreateTokenRequest struct {
	User       int64  `json:"user,omitempty"`
	Org        int64  `json:"org,omitempty"`
	Version    int64  `json:"version,omitempty"`
	TokenType  string `json:"token_type,omitempty"`
	TokenValue string `json:"token_value,omitempty"`
}
