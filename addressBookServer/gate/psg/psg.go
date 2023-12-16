package psg

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Record struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
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

// RecordAdd добавляет новую запись в базу данных.
func (p *Psg) RecordAdd(record Record) (err error) {
	defer func() { err = errors.Wrap(err, "postgres (p *Psg) RecordAdd()") }()
	query := `INSERT INTO address ("firstname", "lastname", "middlename", "phone", "address") VALUES ($1, $2, $3, $4, $5);`
	_, err = p.conn.Exec(context.Background(), query, record.Name, record.LastName, record.MiddleName, record.Phone, record.Address)
	return
}

// RecordUpdate обновляет существующую запись в базе данных по номеру телефона.
func (p *Psg) RecordUpdate(record Record) error {
	_, err := p.conn.Exec(context.Background(), "UPDATE address SET firstname = $1, lastname = $2, middlename = $3, address = $5 WHERE phone = $4", record.Name, record.LastName, record.MiddleName, record.Phone, record.Address)
	return err
}

// RecordsGet возвращает записи из базы данных на основе предоставленных полей Record.
func (p *Psg) RecordGet(record Record) ([]Record, error) {
	var records []Record
	rows, err := p.conn.Query(context.Background(), "SELECT id, firstname, lastname, middlename, phone, address FROM address WHERE firstname = $1 AND lastname = $2 OR middlename = $3 OR phone = $4 OR address = $5", record.Name, record.LastName, record.MiddleName, record.Phone, record.Address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var r Record
		err := rows.Scan(&r.ID, &r.Name, &r.LastName, &r.MiddleName, &r.Phone, &r.Address)
		if err != nil {
			return nil, err
		}
		records = append(records, r)
	}

	fmt.Println(records)
	return records, nil
}

// RecordDeleteByPhone удаляет запись из базы данных по номеру телефона.
func (p *Psg) RecordDeleteByPhone(phone string) error {
	_, err := p.conn.Exec(context.Background(), "DELETE FROM address WHERE phone = $1", phone)
	if err != nil {
		return err
	}
	return nil
}
