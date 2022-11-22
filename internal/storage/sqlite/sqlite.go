package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

type Storage struct {
	db *sql.DB
}

func New(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("can't open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	q := `
	CREATE TABLE IF NOT EXISTS recipients (id INTEGER PRIMARY KEY AUTOINCREMENT, mail_addr TEXT, name TEXT, surname TEXT, birthday TEXT);
	CREATE TABLE IF NOT EXISTS templates (id INTEGER PRIMARY KEY AUTOINCREMENT, template TEXT);
	CREATE TABLE IF NOT EXISTS mailing_tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, exec_time TEXT, mailing_send_id TEXT, mail_addrs TEXT, template_id INTEGER);
	`

	_, err = db.ExecContext(context.TODO(), q)
	if err != nil {
		return nil, fmt.Errorf("can't create table: %w", err)
	}

	return &Storage{
		db: db,
	}, nil
}
