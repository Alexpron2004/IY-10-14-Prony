package stdhttp

import (
	"Addressprj/gate/psg"
	"fmt"
	"net/http"
	"text/template"
)

// Record представляет запись в адресной книге.
type Record struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
}

// Controller обрабатывает HTTP запросы для адресной книги.
type Controller struct {
	DB  *psg.Psg
	Srv *http.Server
}

// NewController создает новый Controller.
func NewController(addr string, db *psg.Psg) *Controller {
	var contr = &Controller{DB: db, Srv: &http.Server{}}
	http.HandleFunc("/", index)
	http.HandleFunc("/create", createpage)
	http.HandleFunc("/get", getpage)
	http.HandleFunc("/update", updatepage)
	http.HandleFunc("/delete", deletepage)
	http.HandleFunc("/RecordAdd", contr.RecordAdd)
	http.HandleFunc("/RecordGet", contr.RecordGet)
	http.HandleFunc("/RecordUpdate", contr.RecordUpdate)
	http.HandleFunc("/RecordDeleteByPhone", contr.RecordDeleteByPhone)
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
	return contr
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func createpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func getpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/get.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "get", nil)
}

func updatepage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/update.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "update", nil)
}

func deletepage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/delete.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "delete", nil)
}

// RecordAdd обрабатывает HTTP запрос для добавления новой записи.
func (c *Controller) RecordAdd(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	lastname := r.FormValue("last_name")
	middleName := r.FormValue("middle_name")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	p := psg.Record{Name: name, LastName: lastname, MiddleName: middleName, Phone: phone, Address: address}
	c.DB.RecordAdd(p)
}

// RecordsGet обрабатывает HTTP запрос для получения записей на основе предоставленных полей Record.
func (c *Controller) RecordGet(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	lastname := r.FormValue("last_name")
	middleName := r.FormValue("middle_name")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	p := psg.Record{Name: name, LastName: lastname, MiddleName: middleName, Phone: phone, Address: address}
	c.DB.RecordGet(p)
}

// RecordUpdate обрабатывает HTTP запрос для обновления записи.
func (c *Controller) RecordUpdate(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	lastname := r.FormValue("last_name")
	middleName := r.FormValue("middle_name")
	phone := r.FormValue("phone")
	address := r.FormValue("address")
	p := psg.Record{Name: name, LastName: lastname, MiddleName: middleName, Phone: phone, Address: address}
	c.DB.RecordUpdate(p)
}

// RecordDeleteByPhone обрабатывает HTTP запрос для удаления записи по номеру телефона.
func (c *Controller) RecordDeleteByPhone(w http.ResponseWriter, r *http.Request) {
	phone := r.FormValue("phone")
	p := psg.Record{Phone: phone}
	c.DB.RecordUpdate(p)
}
