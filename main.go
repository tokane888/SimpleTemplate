package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/now", handleClockTpl)
	http.ListenAndServe(":8080", nil)
}

func handleClockTpl(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("clock.tpl"))
	m := map[string]string{
		"Date": "2019-06-23",
		"Time": "22:12",
	}
	tpl.Execute(w, m)
}
