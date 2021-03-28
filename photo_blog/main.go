package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie := getCookie(w, req)

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("photo")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()

		// Split file extension
		ext := strings.Split(fh.Filename, ".")[1]
		// Generate sha for file
		h := sha1.New()
		io.Copy(h, mf)
		// Build Filename
		filename := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		// Create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "images", filename)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		// Copy
		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// Add filename to cookie
		cookie = appendImage(w, cookie, filename)
	}

	data := strings.Split(cookie.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", data[1:])
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	cookie, err := req.Cookie("session")
	if err != nil {
		sid, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session",
			Value: sid.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func appendImages(w http.ResponseWriter, cookie *http.Cookie) *http.Cookie {
	values := []string{"image1.jpg", "image2.jpg", "image3.jpg"}
	s := cookie.Value

	for _, value := range values {
		if !strings.Contains(s, value) {
			s += "|" + value
		}
	}
	cookie.Value = s
	http.SetCookie(w, cookie)
	return cookie
}

func appendImage(w http.ResponseWriter, cookie *http.Cookie, filename string) *http.Cookie {
	s := cookie.Value
	if !strings.Contains(s, filename) {
		s += "|" + filename
	}
	cookie.Value = s
	http.SetCookie(w, cookie)
	return cookie
}