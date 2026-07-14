package server

import (
	"ascii-web-app/src/ascii"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Text   string
	Banner string
	Result string
}

var tmpl, err = template.ParseFiles("templates/index.html")

func HandleHome(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		w.WriteHeader(404)
		return
	}

	if r.URL.Path == "/" {
		w.WriteHeader(200)
	}
	tmpl.Execute(w, nil)
}
func Handleascii(w http.ResponseWriter, r *http.Request) {
	rezult := PageData{}

	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		return
	}
	if r.Method == "POST" {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		exten := r.FormValue("extensions")
		save := r.FormValue("save")

		if text == "" && banner == "" {
			w.WriteHeader(400)
			return
		}
		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			w.WriteHeader(404)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			return
		}

		baner, _ := ascii.Banner(banner)
		render := ascii.Render(text, baner)

		rezult.Text = text
		rezult.Banner = banner
		rezult.Result = render

		if save == "download" {
			filename := fmt.Sprintf("asciiart.%s", exten)

			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
			if exten == "csv" {
				w.Header().Set("Content-Type", "text/csv")
				lines := strings.Split(render, "\n")
				csv := ""
				for _, line := range lines {
					csv += fmt.Sprintf("\"%s\"\n", line)
				}
				w.Write([]byte(csv))
				return
			} else if exten == "pdf" {
				w.Header().Set("Content-Type", "application/pdf")
				w.Write([]byte(render))
				return
			} else {
				w.Header().Set("Content-Type", "text/plain")
			}
			w.Write([]byte(render))
			return
		}
	}
	tmpl.Execute(w, rezult)
}

func Handleswitch(w http.ResponseWriter, r *http.Request) {
	rezult := PageData{}

	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		return
	}
	if r.Method == "GET" {
		text := r.URL.Query().Get("text")
		banner := r.URL.Query().Get("banner")

		if text == "" && banner == "" {
			w.WriteHeader(400)
			return
		}
		if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
			w.WriteHeader(404)
			return
		}
		if err != nil {
			w.WriteHeader(500)
			return
		}

		baner, _ := ascii.Banner(banner)
		render := ascii.Render(text, baner)

		rezult.Text = text
		rezult.Banner = banner
		rezult.Result = render
	}
	tmpl.Execute(w, rezult)
}
