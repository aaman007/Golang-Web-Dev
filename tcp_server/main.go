package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

type Name struct{
	FirstName string
	LastName string
}

type PersonalInfo struct {
	Institution string
	Country string
}

type ProfessionalInfo struct {
	Profession string
	Company string
}

type ContactInfo struct {
	Website string
	Phone string
}

type User struct {
	Name
	PersonalInfo
	ProfessionalInfo
	ContactInfo
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	li, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]

	fmt.Println("---- METHOD : ", method, "-------")
	fmt.Println("---- URI : ", uri, "-------")

	if method == "GET" {
		if uri == "/" {
			index(conn)
		} else if uri == "/about" {
			about(conn)
		} else if uri == "/contact" {
			contact(conn)
		} else if uri == "/apply" {
			apply(conn)
		}
	} else if method == "POST" {
		if uri == "/apply" {
			applyProcess(conn)
		}
	}
}

func includeHeaders(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func htmlToString(filename string) string {
	data := User{
		Name{ "Amanur", "Rahman" },
		PersonalInfo{ "Metropolitan University", "Bangladesh" },
		ProfessionalInfo{ "Software Engineer", "LIILAB" },
		ContactInfo{ "aaman.sunny007.wordpress.com", "01727272626" },
	}

	var buff bytes.Buffer
	err := tpl.ExecuteTemplate(&buff, filename, data)
	if err != nil {
		log.Fatalln(err)
	}
	return buff.String()
}

func index(conn net.Conn) {
	body := htmlToString("tcp_index.gohtml")
	includeHeaders(conn, body)
}

func about(conn net.Conn) {
	body := htmlToString("tcp_about.gohtml")
	includeHeaders(conn, body)
}

func contact(conn net.Conn) {
	body := htmlToString("tcp_contact.gohtml")
	includeHeaders(conn, body)
}

func apply(conn net.Conn) {
	body := htmlToString("tcp_apply.gohtml")
	includeHeaders(conn, body)
}

func applyProcess(conn net.Conn) {
	body := htmlToString("tcp_apply_process.gohtml")
	includeHeaders(conn, body)
}