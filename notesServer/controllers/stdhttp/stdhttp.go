package stdhttp

import (
	"Nodeprj/gate/psg"
	"fmt"
	"net/http"
	"text/template"
)

// Node представляет запись в адресной книге.

type Node struct {
	ID        int64  `json:"id"` // уникальный индекс ноды. Необходим для того, чтобы можно было удалять ноды из списка
	Title     string `json:"title"`
	Anons     string `json:"anons"`
	Full_text string `json:"full_text"`
}

// Controller обрабатывает HTTP запросы для адресной книги.
type Controller struct {
	DB  *psg.Psg
	Srv *http.Server
}

// NewController создает новый Controller.
func NewController(addr string, db *psg.Psg) *Controller {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", createpage)
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
	return nil
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

// RecordAdd обрабатывает HTTP запрос для добавления новой записи.
func (c *Controller) NodeAdd(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")
	p := psg.Node{Title: title, Anons: anons, Full_text: full_text}
	c.DB.NodeAdd(p)
}
