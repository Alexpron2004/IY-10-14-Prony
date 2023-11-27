package psg

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"text/template"

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

type Cond struct {
	Lop    string
	PgxInd string
	Field  string
	Value  any
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
	query := `INSERT INTO Address ("firstname", "lastname", "middlename", "phone", "address") VALUES ($1, $2, $3, $4, $5);`
	_, err = p.Exec(context.Background(), query, record.firstname, record.LastName, record.MiddleName, record.Phone, record.Address)
	return
}

// RecordsGet возвращает записи из базы данных на основе предоставленных полей Record.
func (p *Psg) RecordsGet(record Record) ( []Record, err error) {
	defer func() { err = errors.Wrap(err, "postgres (p *Psg) RecordsGet()") }()

	sqlFields, values, err := StructToFieldsValues(r, "sql.field")
	if err != nil {
		return
	}

	var conds []Cond

	for i := range sqlFields {
		if i == 0 {
			conds = append(conds, Cond{
				Lop:    "",
				PgxInd: "$" + strconv.Itoa(i+1),
				Field:  sqlFields[i],
				Value:  values[i],
			})
			continue
		}
		conds = append(conds, Cond{
			Lop:    "AND",
			PgxInd: "$" + strconv.Itoa(i+1),
			Field:  sqlFields[i],
			Value:  values[i],
		})
	}

	query := `
	SELECT 
		id, name, last_name, middle_name, address, phone
	FROM
	    address_book
	WHERE
		{{range .}} {{.Lop}} {{.Field}} = {{.PgxInd}}{{end}}
;
`
	tmpl, err := template.New("").Parse(query)
	if err != nil {
		return
	}

	var sb strings.Builder
	err = tmpl.Execute(&sb, conds)
	if err != nil {
		return
	}
	fmt.Println(sb.String())
	return rw.err
}

// RecordUpdate обновляет существующую запись в базе данных по номеру телефона.
func (p *Psg) RecordUpdate(record Record) error {
	defer func() { err = errors.Wrap(err, "postgres (p *Psg) RecordUpdate()") }()
	query := `INSERT INTO test ("firstname", "lastname", "middlename", "phone", "address") VALUES ($1, $2, $3, $4, $5) WHERE "phone"=$4;`
	_, err = p.Exec(context.Background(), query, record.firstname, record.LastName, record.MiddleName, record.Phone, record.Address)
	return nil
}

// RecordDeleteByPhone удаляет запись из базы данных по номеру телефона.
func (p *Psg) RecordDeleteByPhone(phone string) error {
	defer func() { err = errors.Wrap(err, "postgres (p *Psg) RecordDeleteByPhone()") }()
	query := `DELETE * FROM Address WHERE "phone"=$4;`
	_, err = p.conn.Exec(context.Background(), query, record.firstname, record.LastName, record.MiddleName, record.Phone, record.Address)
	return nil
}
