package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
)

func (app *application) startSubscription() {

	sc, err := stan.Connect("test-cluster", "client-123")
	if err != nil {
		app.errorLog.Println(err.Error())
	}

	go app.publish(sc)

	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		app.dataValidation(m.Data)
	}, stan.DurableName("my-durable"))
	if err != nil {
		app.errorLog.Println(err.Error())
	}
}

func (app *application) publish(sc stan.Conn) {
	for i:=1; i<=3; i++ {
		nameFile := fmt.Sprintf("dataFilesForTest(json)/model_%d.json", i)
		b, err := ioutil.ReadFile(nameFile)
		if err != nil {
			app.errorLog.Println(err.Error())
		}
		err = sc.Publish("foo", b)
		if err != nil {
			app.errorLog.Println(err.Error())
		}
	}
}
