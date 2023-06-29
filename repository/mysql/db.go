package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

type MysqlDb struct {
	config Config
	db     *sql.DB
}

func (m *MysqlDb) Conn() *sql.DB {
	return m.db
}

func New(cfg Config) *MysqlDb {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		panic(fmt.Errorf("cant open mysql db: %v", err))
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MysqlDb{db: db, config: cfg}
}
