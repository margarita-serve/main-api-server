package minio

// Config represent Config
type Config struct {
	Endpoint        string `yaml:"code" json:"code"`
	AccessKeyID     string `yaml:"name" json:"name"`
	SecretAccessKey string `yaml:"server" json:"server"`
	UseSSL          bool   `yaml:"enable" json:"enable"`
}
