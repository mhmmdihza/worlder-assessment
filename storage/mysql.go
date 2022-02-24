package storage

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

const driver = "mysql"

func New(config *viper.Viper) ( *sqlx.DB, error) {
	db, err := Connectx(config)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewDBStringFromConfig(config *viper.Viper) (string, error) {
	user := config.GetString("database.user")
	host := config.GetString("database.host")
	port := config.GetString("database.port")
	dbName:= config.GetString("database.dbname")
	pass := config.GetString("database.password")
	return fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true",user,pass,host,port,dbName),nil
}

func Connectx(config *viper.Viper) (*sqlx.DB, error) {
	dbString, err := NewDBStringFromConfig(config)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect(driver, dbString)
	if err != nil {
		return nil, err
	}

	return db, nil
}