package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-calltask/config"
	log "go-calltask/log"
)

var (
	MysqlDBInstance *MysqlDB
	MysqlConfig     = config.DBConfig
)

type MysqlDB struct {
	DB *sql.DB
}

func (mdb *MysqlDB) connect() (error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?autocommit=true&timeout=3s", MysqlConfig.User, MysqlConfig.Passwd, "tcp", MysqlConfig.Host, MysqlConfig.Port, MysqlConfig.Db)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.LOGGER.Error("open mysql failed %s",err.Error())
		return err
	}
	DB.SetConnMaxLifetime(0)
	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(5)
	mdb.DB = DB
	return nil
}

func GetMysqlDb() (*MysqlDB) {
	return MysqlDBInstance
}

func Run() {
	MysqlConfig = config.DBConfig
	mysqlDb := new(MysqlDB)
	err := mysqlDb.connect()
	if err != nil {
		return
	}
	log.LOGGER.Info("%s", "connect mysql success")
	MysqlDBInstance = mysqlDb
}
