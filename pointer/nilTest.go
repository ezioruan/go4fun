package main

import (
    "fmt"
)

type Y struct{
    X int
    Y string
}


func (self *Y) Y1 (){
    fmt.Println("print something:XXX")
    fmt.Println("print X",self.X)
}

func returnNil() *Y  {
    return nil
}


func main() {
    y := &Y{}
    fmt.Println("before set to nil:",y)
    y = returnNil()
    fmt.Println("after set to nil:",y)
    y.Y1()
}
