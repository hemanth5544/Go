package main


import (
	"fmt"
)

func main() {
	//slice	
	fmt.Println("Slices jinxx")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println("Days of the week:", days)

//trditional for loop
	// for	i:=0; i<len(days); i++ {
	// 	fmt.Println("Day", i+1, ":", days[i])
	// }


	// for range loop
for i:= range days {
	fmt.Println("Day", i+1, ":", days[i])
}

}
