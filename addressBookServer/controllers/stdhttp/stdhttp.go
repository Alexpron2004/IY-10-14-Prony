package stdhttp

import (
	"Modprj/gate/psg"
	"fmt"
	"net/http"
	"text/template"
)

// Record представляет запись в адресной книге.
type Record struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name,omitempty"`
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
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/RecordAdd", RecordAdd)
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
	return nil
}

// страница с формой
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

// RecordAdd обрабатывает HTTP запрос для добавления новой записи.
func (c *Controller) RecordAdd(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	http.Redirect(w, r, "/", http.StatusSeeOther)
	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	}
}

// RecordsGet обрабатывает HTTP запрос для получения записей на основе предоставленных полей Record.
func (c *Controller) RecordsGet(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать
}

// RecordUpdate обрабатывает HTTP запрос для обновления записи.
func (c *Controller) RecordUpdate(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать
}

// RecordDeleteByPhone обрабатывает HTTP запрос для удаления записи по номеру телефона.
func (c *Controller) RecordDeleteByPhone(w http.ResponseWriter, r *http.Request) {
	// TODO: Реализовать
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func save_article(w http.ResponseWriter, r *http.Request) {

}

func handlerequest() {

}
