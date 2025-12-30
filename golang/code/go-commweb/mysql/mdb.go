package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-commweb/config"
	log "go-commweb/log"
)

var (
	mdbIns *mysqlDB
	dbConf config.Mysql
)

type mysqlDB struct {
	DB *sql.DB
}

func (mdb *mysqlDB) connect() (error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?autocommit=true&timeout=3s", dbConf.User, dbConf.PassWd, "tcp", dbConf.Host, dbConf.Port, dbConf.Db)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.LOGGER.Error("open mysql failed %v",err)
		return err
	}
	DB.SetConnMaxLifetime(0)
	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(5)
	mdb.DB = DB
	return nil
}

func GetMysqlDb() (*mysqlDB) {
	return mdbIns
}

func Run() {
	dbConf = config.ConfigParam.MysqlConfig
	mysqlDb := new(mysqlDB)
	err := mysqlDb.connect()
	if err != nil {
		return
	}
	log.LOGGER.Info("%s", "connect mysql success")
	mdbIns = mysqlDb
}
