package main

import (
	"Modprj/controllers/stdhttp"
	"Modprj/gate/psg"
	"fmt"
	"net/http"
	"os"
)

func main() {
	s, err := psg.NewPsg("postgres://postgres:1234@localhost:5432/node", "prony", os.Getenv("1234"))
	if err != nil {
		fmt.Println("Error")
	}
	stdhttp.NewController("localhost:8080", s)
	http.ListenAndServe(":8080", nil)
}
