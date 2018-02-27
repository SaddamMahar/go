package main

import("net/http")

func main() {
	serveWeb()
}

func serveWeb() {
	http.HandleFunc("/",serveFunc)
	http.HandleFunc("/test/",serveContact)
	http.ListenAndServe(":8080",nil)
}

func serveFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world."))
}


func serveContact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world. This is dumy page."))
}