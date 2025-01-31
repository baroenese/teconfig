package teconfig

import (
	"os"
	"testing"
)

func TestDefaultPgCOnfig(t *testing.T) {
	expectedHost := "localhost"
	expectedPort := uint(5432)
	expectedUsername := "john"
	expectedPassword := "example"
	expectedSslMode := "disabled"
	expectedDbName := "db_example"
	cfg := DefaultConfig()
	if expectedHost != cfg.DBConfig.Host {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", expectedHost, expectedHost, cfg.DBConfig.Host, cfg.DBConfig.Host)
	}
	if expectedPort != cfg.DBConfig.Port {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", expectedPort, expectedPort, cfg.DBConfig.Port, cfg.DBConfig.Port)
	}
	if expectedUsername != cfg.DBConfig.Username {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", expectedUsername, expectedUsername, cfg.DBConfig.Username, cfg.DBConfig.Username)
	}
	if expectedPassword != cfg.DBConfig.Password {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", expectedPassword, expectedPassword, cfg.DBConfig.Password, cfg.DBConfig.Password)
	}
	if expectedSslMode != cfg.DBConfig.SslMode {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", expectedSslMode, expectedSslMode, cfg.DBConfig.SslMode, cfg.DBConfig.SslMode)
	}
	if expectedDbName != cfg.DBConfig.DBName {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", expectedDbName, expectedDbName, cfg.DBConfig.DBName, cfg.DBConfig.DBName)
	}
}

func TestPgConfig_ConnStr(t *testing.T) {
	cfg := pgConfig{
		Host:     "db.example.com",
		Port:     5432,
		DBName:   "db_example",
		Username: "admin",
		Password: "example",
	}
	expectedConn := "postgres://admin:example@db.example.com:5432/db_example"
	if result := cfg.ConnStr(); result != expectedConn {
		t.Errorf("Expected %s, Got: %s", expectedConn, result)
	}
}

func TestPgConfig_LoadFromEnv_NoChangeDueToValueReceiver(t *testing.T) {
	originalConfig := pgConfig{
		Host:     "original",
		Port:     9876,
		DBName:   "original_db",
		Username: "original_username",
		Password: "original_password",
		SslMode:  "original_ssl",
	}
	cfg := originalConfig
	os.Setenv("KAD_DB_HOST", "2234")
	defer os.Unsetenv("KAD_DB_HOST")
	cfg.LoadFromEnv()
	if originalConfig.Host != cfg.Host {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", originalConfig.Host, originalConfig.Host, cfg.Host, cfg.Host)
	}
	if originalConfig.Port != cfg.Port {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", originalConfig.Port, originalConfig.Port, cfg.Port, cfg.Port)
	}
	if originalConfig.DBName != cfg.DBName {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", originalConfig.DBName, originalConfig.DBName, cfg.DBName, cfg.DBName)
	}
	if originalConfig.Username != cfg.Username {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", originalConfig.Username, originalConfig.Username, cfg.Username, cfg.Username)
	}
	if originalConfig.Password != cfg.Password {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", originalConfig.Password, originalConfig.Password, cfg.Password, cfg.Password)
	}
	if originalConfig.SslMode != cfg.SslMode {
		t.Errorf("Expected: %v (%T), Got: %v (%T)", originalConfig.SslMode, originalConfig.SslMode, cfg.SslMode, cfg.SslMode)
	}
}
