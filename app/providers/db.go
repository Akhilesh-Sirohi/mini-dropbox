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
	DB   *DbProvider
)

type DbProvider struct {
	*gorm.DB
}

type DbConfig struct {
	username string
	password string
	host     string
	port     string
	protocol string
	dbName   string
}

func (db *DbProvider) Setup(cfg DbConfig) error {
	var err error
	once.Do(func() {
		db.DB, err = gorm.Open(mysql.Open(getDatabasePath(cfg)))
		if err != nil {
			logrus.WithError(err).Error("DATABASE_CONNECTION_ERROR")
		}
	})

	return err
}

func getDatabasePath(dbconfig DbConfig) string {
	path := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.username, dbconfig.password, dbconfig.protocol, dbconfig.host, dbconfig.port, dbconfig.dbName)
	return path
}

func (db *DbProvider) Ping() error {
	tx := DB.Exec("SELECT 1 + 1")

	return tx.Error
}
