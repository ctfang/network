package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/byte", sayByte)
	http.ListenAndServe(":8081", nil)
}

func sayByte(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(" say  byte byte!!"))
}
