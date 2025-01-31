package teconfig

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

func loadEnvStr(key string, result *string) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}
	*result = s
}

func loadEnvUint(key string, result *uint) {
	s, ok := os.LookupEnv(key)
	if !ok {
		return
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return
	}
	*result = uint(n)
}

type pgConfig struct {
	Host     string `yaml:"host" json:"host"`
	Port     uint   `yaml:"port" json:"port"`
	DBName   string `yaml:"db_name" json:"db_name"`
	Username string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
	SslMode  string `yaml:"ssl_mode" json:"ssl_mode"`
}

func (pgConfig pgConfig) ConnStr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", pgConfig.Username, pgConfig.Password, pgConfig.Host, pgConfig.Port, pgConfig.DBName)
}

func (pgConfig pgConfig) LoadFromEnv() {
	loadEnvStr("KAD_DB_HOST", &pgConfig.Host)
	loadEnvUint("KAD_DB_PORT", &pgConfig.Port)
	loadEnvStr("KAD_DB_NAME", &pgConfig.DBName)
	loadEnvStr("KAD_DB_USERNAME", &pgConfig.Username)
	loadEnvStr("KAD_DB_PASSWORD", &pgConfig.Password)
	loadEnvStr("KAD_DB_SSL", &pgConfig.SslMode)
}

func defaultPgConfig() pgConfig {
	return pgConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "john",
		Password: "example",
		SslMode:  "disabled",
		DBName:   "db_example",
	}
}

type listenConfig struct {
	Host string `yaml:"host" json:"host"`
	Port uint   `yaml:"port" json:"port"`
}

func (listenConfig listenConfig) Addr() string {
	return fmt.Sprintf("%s:%d", listenConfig.Host, listenConfig.Port)
}

func (listenConfig *listenConfig) loadFromEnv() {
	loadEnvStr("KAD_LISTEN_HOST", &listenConfig.Host)
	loadEnvUint("KAD_LISTEN_PORT", &listenConfig.Port)
}

func defaultListenConfig() listenConfig {
	return listenConfig{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

type TeConfig struct {
	Listen   listenConfig `yaml:"listen" json:"listen"`
	DBConfig pgConfig     `yaml:"db" json:"db"`
}

func (c *TeConfig) LoadFromEnv() {
	c.Listen.loadFromEnv()
}

func loadConfigFromReader(r io.Reader, c *TeConfig) error {
	return yaml.NewDecoder(r).Decode(c)
}

func LoadConfigFromFile(fn string, c *TeConfig) error {
	_, err := os.Stat(fn)
	if err != nil {
		return err
	}
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()
	return loadConfigFromReader(f, c)
}
