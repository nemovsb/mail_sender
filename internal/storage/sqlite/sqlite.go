package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"mail_sender/internal/app"
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

// Create recipient
func (s *Storage) CreateRecipients(recipients []app.Recipient) (rowsCount uint, err error) {

	insertQuery, err := s.db.Prepare(`INSERT INTO recipients (mail_addr, name, surname, birthday) VALUES (?, ?, ?, ?);`)

	tx, err := s.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("can't create recipients: %w", err)
	}

	for _, recp := range recipients {
		res, err := tx.Stmt(insertQuery).Exec(recp.MailAddr, recp.Name, recp.Surname, recp.Birthday)
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("can't create recipient (mail = %s): %w", recp.MailAddr, err)
		}
		rows, err := res.RowsAffected()
		rowsCount += uint(rows)
	}
	err = tx.Commit()
	if err != nil {
		return 0, fmt.Errorf("can't create recipients: %w", err)
	}
	return rowsCount, err

}

// // Get recipients by email-adresses
// func (s *Storage) GetRecipients(mailAddrs []string) ([]app.Recipient, error) {

// }

// // Get all recipients from storage
// func (s *Storage) GetAllRecipients() []app.Recipient {

// }

// Create template
func (s *Storage) CreateTemplate(templ string) (uint, error) {

	q := `
	INSERT INTO templates (template) 
	VALUES (?);`

	res, err := s.db.ExecContext(context.TODO(), q, templ)
	if err != nil {
		return 0, fmt.Errorf("can't save template: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("can't get template ID: %w", err)
	}

	return uint(id), nil

}

// Get template by id
func (s *Storage) GetTemplate(id uint) (templ string, err error) {

	q := `
	SELECT template 
	FROM templates 
	WHERE templates.id = ?;`

	err = s.db.QueryRowContext(context.TODO(), q, id).Scan(&templ)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("template with id = %d does not exist : %w", id, err)
	}
	if err != nil {
		return "", fmt.Errorf("can't get template : %w", err)
	}

	return templ, nil
}

// Get all templates from storage
func (s *Storage) GetAllTemplates() (templs []string, err error) {

	q := `
	SELECT id, template 
	FROM templates;`

	err = s.db.QueryRowContext(context.TODO(), q).Scan(&templs)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no templates : %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("can't get templates : %w", err)
	}

	return templs, nil
}

// func (s *Storage) AddMailingTask(app.MailingTask) string {

// }

// func (s *Storage) GetMailingTasks() []app.MailingTask {

// }