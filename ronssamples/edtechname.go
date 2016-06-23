package main

import (

	"fmt"
	"math/rand"
	"time"
)

var wantnew string
var newword string
var a = []string{"Ed", "Surge", "Learn", "Open", "Lesson", "Bright", "Course", "Class", "Teacher", "Chalk", "Board", "Backpack", "Classroom", "Hub", "Book", "Scholar", "Smarter", "Hooli"}

func Shuffle(a []string) {
    for i := range a {
        j := rand.Intn(i + 1)
        a[i], a[j] = a[j], a[i]
    }
}

func Asknew (string) {

	fmt.Print("OK, please let me know a word you want to mix into the randomizer: ")
	fmt.Scan(&newword)
	a = append([]string{newword}, a...)
	return

}

func main() {
    
    rand.Seed(time.Now().UnixNano())
    Shuffle(a)
 	fmt.Print("Do you want to add another word into the company name randomizer? ")
 	fmt.Scan(&wantnew)
 	if wantnew == "yes" {
 	Asknew(wantnew)
	}
 	
    b:=a[0]
    c:=a[1]
    fmt.Println("Great, Your New Company Name is: ",b+c)

}