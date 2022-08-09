package config

// Connectors represent available connectors
type Connectors struct {
	Identity       Identity        `json:"identity" yaml:"identity"`
	Storages       Storages        `json:"storages" yaml:"storages"`
	DriftServer    DataDriftServer `json:"driftServer" yaml:"driftServer"`
	AccuracyServer AccuracyServer  `json:"accuracyServer" yaml:"accuracyServer"`
	GraphServer    GraphServer     `json:"graphServer" yaml:"graphServer"`
	Kafka          KafkaServer     `json:"Kafka" yaml:"kafka"`
}

// Identity type
type Identity struct {
	EA2M EA2M `json:"ea2m" yaml:"ea2m"`
}

// EA2M Type
type EA2M struct {
	Server            string `json:"server" yaml:"server"`
	ClientAccessKey   string `json:"clientAccessKey" yaml:"clientAccessKey"`
	ClientSecretKey   string `json:"clientSecretKey" yaml:"clientSecretKey"`
	AllowDevToken     bool   `json:"allowDevToken" yaml:"allowDevToken"`
	DevIdentityToken  string `json:"devIdentityToken" yaml:"devIdentityToken"`
	DevIdentityClaims string `json:"devIdentityClaims" yaml:"devIdentityClaims"`
}

type Storages struct {
	Minio Minio `json:"storage" yaml:"storage"`
}

type Minio struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyID     string `json:"accessKeyID" yaml:"accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey" yaml:"secretAccessKey"`
	UseSSL          bool   `json:"useSSL" yaml:"useSSL"`
}

type DataDriftServer struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

type AccuracyServer struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

type GraphServer struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

type KafkaServer struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	GroupID         string `json:"groupID"`
	AutoOffsetReset string `json:"autoOffsetReset"`
}
