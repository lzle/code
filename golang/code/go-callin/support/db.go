package support

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go-callin/core"
)

type Mysql struct {
	DB *sql.DB
}

func (m *Mysql) Connect() (error) {
	dsn := core.CONFIG.MysqlUrl()

	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		core.LOGGER.Error("open mysql failed " + err.Error())
		return err
	}
	DB.SetConnMaxLifetime(0)
	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(5)
	m.DB = DB
	return nil
}

func (m *Mysql) GetDB() *sql.DB{
	return m.DB
}


func (m *Mysql) Init() {
	err := m.Connect()
	if err != nil {
		return
	}
	core.LOGGER.Info("connet mysql success")
	core.MYSQL = m
}