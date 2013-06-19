package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
    "log"
)



type P struct {
    X,Y,Z int
    Name string
}

type Q struct {
    X,Y *int32
    Name string
}

func main() {
    var network bytes.Buffer
    encoder := gob.NewEncoder(&network)
    decoder := gob.NewDecoder(&network)

    err := encoder.Encode(P{1,2,3,"ezio"})
    if err != nil {
        log.Fatal("ecode error",err)
    }

    var q Q
    err = decoder.Decode(&q)
    if err != nil{
        log.Fatal("decode error:",err)
    }
    fmt.Println(q)
    fmt.Printf("%q :{%d,%d}\n",q.Name,*q.X,*q.Y)
}
