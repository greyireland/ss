package main

import (
	"net/http"
	"flag"
	"fmt"
)

func main() {
	dir := flag.String("d", "./", "static server dir ")
	port := flag.String("p", "9090", "static server port ")
	flag.Parse()
	fmt.Printf("static dir [%s],listen port [%s]\n", *dir, *port)
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	err := http.ListenAndServe("0.0.0.0:" + *port, nil)
	if err != nil {
		fmt.Printf("port [%s] is used", *port)
	}
}
