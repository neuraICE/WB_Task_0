package main

import (
	"WB_Task_1/pkg/postgres"
	"log"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

var cache = map[string][]byte{}

func (app application)caching()  {
	err := postgres.SelectInDB()
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	for i := range postgres.M{
		cache[i] = postgres.M[i]
	}
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	app.caching()

	go app.startSubscription()

	app.start()
}
