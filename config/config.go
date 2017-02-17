package config

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

const PolicyDefault = "default"
const PolicyKeepMostUsed = "keep-most-used"

type Config struct {
	Cache  CacheConfig  `json:"Cache"`
	Server ServerConfig `json:"Server"`
}

type CacheConfig struct {
	MaxEntries int    `json:"MaxEntries"`
	MinTTL     int    `json:"MinTTL"`
	Policy     string `json:"Policy"`
}

type ServerConfig struct {
	Address net.IP   `json:"Address"`
	Port    int      `json:"Port"`
	Servers []net.IP `json:"Servers"`
}

func (c *Config) Valid() bool {
	if c.Cache.Policy != PolicyDefault && c.Cache.Policy != PolicyKeepMostUsed {
		return false
	}

	return true
}

func Load() (*Config, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Printf("error opening config file %s", err)
		return nil, err
	}

	decoder := json.NewDecoder(file)
	config := new(Config)
	err = decoder.Decode(config)
	if err != nil {
		log.Printf("error decoding json config: %s", err)
		return nil, err
	}

	if !config.Valid() {
		log.Printf("invalid config")
		return nil, fmt.Errorf("invalid config")
	}

	return config, nil
}

func (c *Config) Store() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Printf("error opening config file")
		return
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&c)
	if err != nil {
		log.Printf("config save err: %s", err)
	}
}
