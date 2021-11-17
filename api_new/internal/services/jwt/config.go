package jwt

type Config struct {
	PrivateKey string `required:"true" yaml:"private_key" env:"REAL_ESTATE_PRIVATE_KEY"`
	KeyID      string `required:"true" yaml:"key_id"`
	Expiration int    `required:"true" yaml:"expiration"`
	Issuer     string `required:"true" yaml:"issuer"`
	SigningAlg string `required:"true" yaml:"singing_alg"`
}
