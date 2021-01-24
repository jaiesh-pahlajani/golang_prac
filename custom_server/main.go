package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	method, url := request(conn)

	// write response
	respond(conn, method, url)
}

func request(conn net.Conn) (string, string) {
	i := 0
	method := ""
	url := ""
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			method = strings.Fields(ln)[0]
			url = strings.Fields(ln)[1]
			fmt.Println("***METHOD", method)
			fmt.Println("***URL", url)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return method, url
}

func respond(conn net.Conn, method string, url string) {

	if method == "GET" {
		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Get request %v</strong></body></html>`
		body = fmt.Sprintf(body, url)
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	} else if method == "POST" {
		body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>POST request %v</strong></body></html>`
		body = fmt.Sprintf(body, url)
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: % d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}

}
