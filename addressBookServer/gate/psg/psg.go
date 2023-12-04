package psg

import (
	"context"
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
func (p *Psg) RecordAdd(record Record) (int64, error) {
	var id int64
	err := p.conn.QueryRow(context.Background(), "INSERT INTO records (name, last_name, middle_name, phone, address) VALUES ($1, $2, $3, $4, $5) RETURNING id", record.Name, record.LastName, record.MiddleName, record.Phone, record.Address).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// RecordsGet возвращает записи из базы данных на основе предоставленных полей Record.
func (p *Psg) RecordsGet(record Record) ([]Record, error) {
	var records []Record
	rows, err := p.conn.Query(context.Background(), "SELECT id, name, last_name, middle_name, phone, address FROM records WHERE name = $1 AND last_name = $2", record.Name, record.LastName)
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
	return records, nil
}

// RecordUpdate обновляет существующую запись в базе данных по номеру телефона.
func (p *Psg) RecordUpdate(record Record) error {
	_, err := p.conn.Exec(context.Background(), "UPDATE records SET name = $1, last_name = $2, middle_name = $3, address = $4 WHERE phone = $5", record.Name, record.LastName, record.MiddleName, record.Address, record.Phone)
	return err
}

// RecordDeleteByPhone удаляет запись из базы данных по номеру телефона.
func (p *Psg) RecordDeleteByPhone(phone string) error {
	commandTag, err := p.conn.Exec(context.Background(), "DELETE FROM records WHERE phone = $1", phone)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("Record not found")
	}
	return nil
}
