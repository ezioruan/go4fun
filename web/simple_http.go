package main

import (
    "net/http"
    "fmt"
)


func HelloServer(writer http.ResponseWriter,request *http.Request){
    writer.Write([]byte("so here you are!"))
}

func main() {
    http.HandleFunc("/hello",HelloServer)
    err := http.ListenAndServe(":8083",nil)
    if err != nil{
        fmt.Println("err",err)
    }
}
