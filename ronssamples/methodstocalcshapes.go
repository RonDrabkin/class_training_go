package main

import "fmt"



type Rectangle struct {
	side1, side2 int
}



func (r Rectangle) perimeter() int {
	return r.side1 + r.side2
}

//func area

func main() {


	fmt.Print("What is side 1? ")
	fmt.Scan(&side1)
	fmt.Print("What is side 2? ")
	fmt.Scan(&side2)

   r := Rectangle{side1, side2}

	fmt.Println(r.perimeter())
}
