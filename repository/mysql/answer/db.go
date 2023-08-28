package mysqlanswer

import "game-app/repository/mysql"

type DB struct {
	conn *mysql.MysqlDb
}

func New(conn *mysql.MysqlDb) *DB {
	return &DB{conn: conn}
}
