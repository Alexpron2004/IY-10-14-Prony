package psg

import (
	"context"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Node struct {
	ID        int64  `json:"id"` // уникальный индекс ноды. Необходим для того, чтобы можно было удалять ноды из списка
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	Full_text string `json:"full_text"`
}

// Psg представляет гейт к базе данных PostgreSQL.
type Psg struct {
	conn *pgxpool.Pool
}

// NewPsg создает новый экземпляр Psg.
func NewPsg(dburl string, login, pass string) (psg *Psg, err error) {
	defer func() { err = errors.Wrap(err, "postgres NewPsg()") }()

	psg = &Psg{}
	psg.conn, err = parseConnectionString(dburl, login, pass)
	if err != nil {
		return nil, err
	}

	err = psg.conn.Ping(context.Background())
	if err != nil {
		err = errors.Wrap(err, "psg.conn.Ping(context.Background())")
		return nil, err
	}

	return

}

func parseConnectionString(dburl, user, password string) (db *pgxpool.Pool, err error) {
	var u *url.URL
	if u, err = url.Parse(dburl); err != nil {
		return nil, errors.Wrap(err, "ошибка парсинга url строки")
	}
	u.User = url.UserPassword(user, password)
	db, err = pgxpool.New(context.Background(), u.String())
	if err != nil {
		return nil, errors.Wrap(err, "ошибка соединения с базой данных")
	}
	return
}

// NodeAdd добавляет новую запись в базу данных.
func (p *Psg) NodeAdd(node Node) (err error) {
	defer func() { err = errors.Wrap(err, "postgres (p *Psg) RecordAdd()") }()
	query := `INSERT INTO node ("title", "anons", "full_text") VALUES ($1, $2, $3);`
	_, err = p.conn.Exec(context.Background(), query, node.Title, node.Anons, node.Full_text)
	return
}

// NodeUpdate обновляет существующую запись в базе данных по номеру телефона.
func (p *Psg) NodeUpdate(node Node) error {
	_, err := p.conn.Exec(context.Background(), "UPDATE node SET title = $1, anons = $2, full_text = $3 WHERE title = $1", node.Title, node.Anons, node.Full_text)
	return err
}
