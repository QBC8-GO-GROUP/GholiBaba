package config

type NatsConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Subject string `json:"subject"`
}
