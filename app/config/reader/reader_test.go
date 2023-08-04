package reader

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	Title string
	Db    TestDbConfig
	Raven TestRavenConfig
}

type TestDbConfig struct {
	Dialect               string
	Protocol              string
	Host                  string
	Port                  int
	Username              string
	Password              string
	SslMode               string
	Name                  string
	MaxOpenConnections    int
	MaxIdleConnections    int
	ConnectionMaxLifetime time.Duration
}

type TestRavenConfig struct {
	ClientName string
	Host       string
	Auth       TestAuthConfig
}

type TestAuthConfig struct {
	Username string
	Password string
}

func TestLoadConfig(t *testing.T) {
	var c TestConfig

	key := strings.ToUpper("cyber_helpdesk") + "_DB_PASSWORD"
	os.Setenv(key, "envpass")

	os.Setenv(strings.ToUpper("cyber_helpdesk")+"_RAVEN_AUTH_PASSWORD", "RAVEN123")

	err := NewConfig(NewOptions("toml", "./testdata", "default")).Load("drone", &c)
	assert.Nil(t, err)
	// Asserts that default value exists.
	assert.Equal(t, "mysql", c.Db.Dialect)
	// Asserts that application environment specific value got overridden.
	assert.Equal(t, 10, c.Db.MaxOpenConnections)
	// Asserts that environment variable was honored.
	assert.Equal(t, "envpass", c.Db.Password)

	//assert that Raven Password was picked from auth
	assert.Equal(t, "RAVEN123", c.Raven.Auth.Password)
}
