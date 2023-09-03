package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IfconfigMeRespData struct {
	IPAddr     string `json:"ip_addr"`
	RemoteHost string `json:"remote_host"`
	UserAgent  string `json:"user_agent"`
	Port       string `json:"port"`
	Method     string `json:"method"`
	Encoding   string `json:"encoding"`
	Mime       string `json:"mime"`
	Via        string `json:"via"`
	Forwarded  string `json:"forwarded"`
}

func main() {
	resp, err := http.Get("https://ifconfig.me/all.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var respData IfconfigMeRespData
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		log.Fatalln(err)
	}
	log.Println(respData)

}
