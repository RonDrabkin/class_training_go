package main

import "fmt"

type Rectangle struct {
    length, width int
}

func (myrect Rectangle) Area() int {
    return myrect.length * myrect.width
}

func main(){

myrect:=Rectangle {4,3}

fmt.Println(myrect)
fmt.Println(myrect.Area())
}