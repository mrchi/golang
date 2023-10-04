package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	lua "github.com/yuin/gopher-lua"
)

const (
	LuaHTTPTypeName = "http"
	PLUGINS_DIR     = "../plugins/"
)

func head(l *lua.LState) int {
	host := l.CheckString(1)
	port := l.CheckInt64(2)
	path := l.CheckString(3)
	reqUrl := fmt.Sprintf("http://%s:%d/%s", host, port, path)
	if resp, err := http.Head(reqUrl); err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LBool(false))
		l.Push(lua.LString(fmt.Sprintf("Request failed: %s", err)))
	} else {
		l.Push(lua.LNumber(resp.StatusCode))
		l.Push(lua.LBool(resp.Header.Get("WWW-Authenticate") != ""))
		l.Push(lua.LString(""))
	}
	return 3
}

func get(l *lua.LState) int {
	host := l.CheckString(1)
	port := l.CheckInt64(2)
	username := l.CheckString(3)
	password := l.CheckString(4)
	path := l.CheckString(5)

	reqUrl := fmt.Sprintf("http://%s:%d/%s", host, port, path)
	client := http.Client{}
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LBool(false))
		l.Push(lua.LString(fmt.Sprintf("Unable to build GET request: %s", err)))
		return 3
	}

	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	resp, err := client.Do(req)
	if err != nil {
		l.Push(lua.LNumber(0))
		l.Push(lua.LBool(false))
		l.Push(lua.LString(fmt.Sprintf("Request failed: %s", err)))
		return 3
	}

	l.Push(lua.LNumber(resp.StatusCode))
	l.Push(lua.LBool(false))
	l.Push(lua.LString(""))
	return 3
}

func register(l *lua.LState) {
	mt := l.NewTypeMetatable(LuaHTTPTypeName)
	l.SetGlobal(LuaHTTPTypeName, mt)
	l.SetField(mt, "head", l.NewFunction(head))
	l.SetField(mt, "get", l.NewFunction(get))
}

func main() {
	l := lua.NewState()
	defer l.Close()

	register(l)
	files, err := os.ReadDir(PLUGINS_DIR)
	if err != nil {
		log.Fatalln(err)
	}
	if len(files) == 0 {
		log.Fatalln("No plugins.")
	}

	for _, file := range files {
		log.Printf("Found plugin: %s\n", file.Name())
		if err := l.DoFile(path.Join(PLUGINS_DIR, file.Name())); err != nil {
			log.Fatalln(err)
		}
	}
}
