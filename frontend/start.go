package main

import (
	"net/http"
	"spider/frontend/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler("frontend/view/template.html"))
	err := http.ListenAndServe(":9300", nil)
	if err != nil {
		panic(err)
	}

}
