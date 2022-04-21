package main

import (
	"flag"
	"net/http"
)



func (app *application) start() {

	addr := flag.String("addr", ":4000", "Сетевой адрес http")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.homePage)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: app.errorLog,
		Handler:  mux,
	}

	app.infoLog.Printf("Запуск сервера на %s", *addr)

	err := srv.ListenAndServe()

	app.errorLog.Fatal(err)
}
