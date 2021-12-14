package main

import "net/http"

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BAR CALLED"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Foo Called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
