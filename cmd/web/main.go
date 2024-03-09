package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type appplication struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "htto network address")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &appplication{
		errorLog: errorlog,
		infoLog:  infolog,
	}

	

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler: app.routes(),
	}

	infolog.Printf("Servidor Funcionando en el Puerto %s", *addr)
	err := srv.ListenAndServe()
	errorlog.Fatal(err)
}
