package main

import (
	"net/http"
)

func (app application) homePage(w http.ResponseWriter, r *http.Request)  {

	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			app.errorLog.Println(err.Error())
		}

		dataId := r.FormValue("dataId")

		app.infoLog.Println(dataId)

		w.Write(cache[dataId])

	}else{
		http.ServeFile(w,r, "./ui/html/home.page.tmpl")
	}
}

//	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
//	if err != nil {
//		app.errorLog.Println(err.Error())
//		http.Error(w, "Internal server Error", 505)
//		return
//	}
//
//	err = ts.Execute(w, nil)
//	if err != nil {
//		app.errorLog.Println(err.Error())
//		http.Error(w, "Internal server Error", 505)
//		return
//	}
//
//}
