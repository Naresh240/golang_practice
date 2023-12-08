package main
import "fmt"

func main() {

	// Difference between varaible and pointer

	// var instituteName string = "NareshLabs"
	// fmt.Println(instituteName)
	// fmt.Println(&instituteName)

	var instituteName string
	fmt.Println("please enter institute name")
	fmt.Scan(&instituteName)


	fmt.Printf("Welcome to %v to learn Devops..!!! \n", instituteName)
}
