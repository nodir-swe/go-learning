//package main

//import "fmt"

//	func main() {
//		fmt.Println("Hello world!")
//		fmt.Println("My name is Nodir")
//	}
package main

import "fmt"

func main() {
	var name string = "Nodir Kuvondik"
	var age int = 16

	targetCountry := "Poland"
	targetYear := 2028
	dailyHour := 7

	fmt.Println("Developer: ", name)
	fmt.Println("Age: ", age)
	fmt.Println("Goal:", targetCountry, "in", targetYear)
	fmt.Println("Daily study hours: ", dailyHour)
}
