package main

import (
	"fmt"
	"math"
	// "math/rand"
	"time"
	// "math"
)

func add(x string ,y string) string{ 
    return x+y
}
func addI(x  ,y int) int{ 
    return x+y
}
type vertex struct{
    x int
    y int

} 

//methods asiifnin to structs
func (v vertex) abs()int{
    return v.x+v.y

}

var anonymus = func (x,y float64) float64  {
    return math.Sqrt(x*x + y*y)
    
}



func split(sum int) (x,y int){// the return x and y vars are inltialsed here only 
    x=sum*4/9
    y=sum-x
    return //no need to retun x y it was above instialsed nad declared

}

//best way

func best(sum int)(int,int){
    x:= sum+10
    fmt.Println(x)
    y:= sum+10
    fmt.Println(y)

    return x ,y
}

const Pi=3.14

type Abser interface{
    Abs() float64
}

type MyFloat float64

type MyError struct{
    When time.Time
    What string
}


func (e *MyError) Error() string {
    return fmt.Sprintf("at %v, %s",
e.When,e.What)
    
}

func run () error {
    return &MyError{
        time.Now(),
        "didnt work",
    }
}

func say(s string)  {
   for i:=0;i<5;i++ {
    time.Sleep(100 *time.Millisecond)
    fmt.Println(s)

   }
    
}

func sums(s []int,c chan int){
    sum:=0

    for _,v := range s{
        sum+=v

    }
    c <- sum    
}

func main() {

    s:= []int {1,2,3}
    c:= make(chan int)
   go sums(s,c)
   go sums(s[len(s)/2:],c)
   x,y:= <-c,<-c
   fmt.Println(x,y,"through channel")

    go say("hellooo gooo")
    say("without gooooo")

    if err:= run(); err !=nil{
        fmt.Println(err)
    }
    // fmt.Println("Hello, 世界");
    // rand.Seed(time.Now().UnixNano())
    // fmt.Println("This a random num from :", rand.Intn(100))
    // fmt.Println(math.Pi)
    // fmt.Println(add("jj","h"))
    // fmt.Println(addI(20,22))

    // a,b:=swap("hello","world")
    // fmt.Println(a,b)
    // fmt.Println(swap("jj","h"))
    // fmt.Println(split(999))
    // fmt.Println(best(111))

    // fmt.Println(Pi)
//    defer fmt.Println("world") //while we add defer it will be in stack and all the remaning code works with flow
//    defer fmt.Println("anothet")


// v:=vertex{1,2}
// fmt.Print((v.x))
// fmt.Print((v.y))
var m= make(map[string] int)
m["answer"]=31
fmt.Println(anonymus(2,2))


for key,value:= range m{
    fmt.Println(key,value)
}
    fmt.Println("Hello")
    fmt.Println("Hello")
    defer fmt.Println("Hello")
    for i:=0;i<5;i++{
        defer fmt.Println(i)
    }
   
    names:=[4] string{
        "Hemanth",
        "yaswanth",
        "harsha",
        "someother",
    }
    // fmt.Println(names)
    a:=names[2:4]
    b:= names[0:2]
    fmt.Println(a,b)
    b[0]="XX"//we are cahnging in slice but the internally slice is pointing to the names arr it will update their alos "XX"
    fmt.Println(a,b)
    fmt.Println(names)


// var sum int =0
sum:=0
for i:=0;i<10;i++{
    sum+=i
}

for sum<1000{
    sum+=sum
    
}


// fmt.Println(sum)


slice:= make([]int, 5,10)
fmt.Println(slice,len(slice),cap(slice),"()()()()")

zuii:= make([]int, 0,8)
PrintSlice("zuii",zuii)


}

func PrintSlice(s string, e []int){
    fmt.Printf(s,len(e),cap(e),"%s len=%d cap=%d")


}



