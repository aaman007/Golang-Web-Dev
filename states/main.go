package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Routes
	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/start", startCounter)
	http.HandleFunc("/increment", incrementCounter)
	http.HandleFunc("/decrement", decrementCounter)
	http.HandleFunc("/remove", removeCounter)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		fmt.Fprintln(w, `<a href="/start">Start Counter</a>`)
		return
	} else if err != nil {
		log.Fatalln(err)
	}
	counter := cookie.Value
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h2>Count Now : "+counter+"</h2>")
	fmt.Fprintln(w, `<a href="/increment">Increment Counter</a><br/>`)
	fmt.Fprintln(w, `<a href="/decrement">Decrement Counter</a><br/>`)
	fmt.Fprintln(w, `<a href="/remove">Remove Counter</a>`)
}

func home(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func startCounter(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("counter")
	if err == nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "counter",
		Value: "0",
	})
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func incrementCounter(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		http.Error(w, "Counter is not initialized", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Fatalln(err)
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func decrementCounter(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err == http.ErrNoCookie {
		http.Error(w, "Counter is not initialized", http.StatusBadRequest)
		return
	} else if err != nil {
		log.Fatalln(err)
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count--
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func removeCounter(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	if err != nil {
		http.Error(w, "Counter is not initialized", http.StatusBadRequest)
		return
	}

	cookie.MaxAge = -1 // deletes cookie

	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
