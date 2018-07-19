package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello geronimo")
	fmt.Println("dfasdfasdf",r.URL.Path)
}
type Message struct{
	Text string
}
func about(w http.ResponseWriter, r *http.Request){
	m:=Message{"welcome to my world"}
	b, err:= json.Marshal(m)
	fmt.Println("dafsd",r.URL.Path)
	if err != nil{
		panic(err)
	}
	w.Write(b)
}
func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080",nil)
	
}