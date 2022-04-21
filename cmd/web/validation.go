package main

import (
	"WB_Task_1/pkg/postgres"
	"encoding/json"
)

func (app *application) dataValidation(b []byte) {
	var model Model
	err := json.Unmarshal(b, &model)
	if err != nil {
		switch e := err.(type) {
		case *json.UnmarshalTypeError:
			app.infoLog.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)
		case *json.InvalidUnmarshalError:
			app.infoLog.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			app.infoLog.Printf(err.Error())
		}
	} else {
		app.dataSave(model.OrderUid, b)
	}
}

func (app *application) dataSave(OrderUid string, b []byte){

	cache[OrderUid] = b

	err := postgres.InsertInDB(OrderUid, b)
	if err != nil {
		app.errorLog.Println(err.Error())
	}
}
