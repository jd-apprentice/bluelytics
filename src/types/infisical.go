package types

type AcessToken struct {
	AccessToken string
	ExpiresIn   int
	TokenType   string
}

type Secret struct {
	ID          string `json:"id"`
	SecretKey   string `json:"secretKey"`
	SecretValue string `json:"secretValue"`
}

type Response struct {
	Secrets []Secret `json:"secrets"`
}
