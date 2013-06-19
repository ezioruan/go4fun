package main

import (
    "fmt"
)

type Y struct{
    X int
    Y string
}


func (self *Y) Y1 (){
    fmt.Println("XXX")
}

func returnNil() *Y  {
    return nil
}


func main() {
    y := &Y{}
    y = returnNil()
    fmt.Println(y)
    y.Y1()
}
