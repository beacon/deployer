package config

import (
	"fmt"
	"io/ioutil"

	"github.com/go-playground/validator"
	"sigs.k8s.io/yaml"
)

type Config struct {
	Mode string `json:"mode,omitempty" validate:"oneof=server worker"`
	Addr string `json:"addr,omitempty"`

	TLS *TLSConfig `json:"tls,omitempty"`
}

type TLSConfig struct {
	CertFile string
	KeyFile  string
}

func New(file string) (*Config, error) {
	// Provide default values
	cfg := &Config{
		Addr: ":9000",
	}
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s:%v", file, err)
	}
	if err := yaml.Unmarshal(raw, cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file %s:%v", file, err)
	}
	return cfg, nil
}

func Validate(cfg *Config) error {
	return validator.New().Struct(cfg)
}
