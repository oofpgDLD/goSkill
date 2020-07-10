package api

import "net/http"

func New() http.Handler{
	mux := http.NewServeMux()
	mux.HandleFunc("/net/test1", Net1)
	return mux
}

func Net1(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("net test1 ok"))
}
