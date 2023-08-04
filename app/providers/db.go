package providers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	once sync.Once
	DB   DbProvider
)

type DbProvider struct {
	*gorm.DB
}

type DbConfig struct {
	Dialect  string
	Protocol string
	Host     string
	Port     int
	Username string
	Password string
	SslMode  string
	Name     string
}

func (db *DbProvider) Setup(cfg DbConfig) error {
	var err error
	once.Do(func() {
		db.DB, err = gorm.Open(mysql.Open(getDatabasePath(cfg)), &gorm.Config{})
		if err != nil {
			logrus.WithError(err).Error("DATABASE_CONNECTION_ERROR")
		}
	})

	return err
}

func getDatabasePath(config DbConfig) string {
	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	return path
}

func (db *DbProvider) Ping() error {
	tx := DB.Exec("SELECT 1 + 1")

	return tx.Error
}
