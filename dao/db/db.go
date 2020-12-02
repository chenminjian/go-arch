package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	mutex sync.RWMutex
	db    *sql.DB
	stmts map[string]*sql.Stmt
}

func New(conf *Conf) (*DB, error) {
	sqlDB, err := initDB(conf)
	if err != nil {
		return nil, err
	}

	db := &DB{
		db:    sqlDB,
		stmts: make(map[string]*sql.Stmt),
	}
	return db, nil
}

func initDB(conf *Conf) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", conf.Username, conf.Password, conf.Host, conf.Port, conf.Name)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(16)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if stmt, ok := db.stmts[query]; ok {
		return stmt, nil
	}

	stmt, err := db.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	db.stmts[query] = stmt
	return stmt, nil
}
