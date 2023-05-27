package data

import (
	"os"
	"log"
	"database/sql"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

const (
	_DB_FILE = "./biletojy.db"
)

type Dao struct {
	db *sql.DB
}

type Ticket struct {
	Id int64
	Title string
	Content string
	Tags string
	CreatedBy string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewDao() *Dao {
	_, err := os.Stat(_DB_FILE)
	exists := err == nil
	db, err := sql.Open("sqlite3", _DB_FILE)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		if _, err := db.Exec(_SQL_INIT); err != nil {
			log.Fatal(err)
		}
	}
	return &Dao{ db: db }
}

func (dao *Dao) Close() {
	dao.db.Close()
	dao.db = nil
}

func AddTicket(dao *Dao, ticket *Ticket) {
	tx, err := dao.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	title_2gram := ticket.Title // TODO bi-gram
	content_2gram := ticket.Content // TODO bi-gram
	comments_2gram := ""
	res, err := tx.Exec(_SQL_ADD_TICKET, ticket.Title, ticket.Content, ticket.Tags, ticket.CreatedBy, ticket.CreatedAt, ticket.UpdatedAt)
	if err != nil {
		log.Fatal(err)
	}
	ticket_id, _ := res.LastInsertId()
	ticket.Id = ticket_id
	tx.Exec(_SQL_ADD_TICKET_HISTORY, ticket.Id, ticket.Title, ticket.Content, ticket.Tags, ticket.CreatedBy, ticket.CreatedAt)
	if _, err := tx.Exec(_SQL_ADD_TICKET_FTS, ticket.Id, title_2gram, content_2gram, ticket.Tags, comments_2gram); err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}
