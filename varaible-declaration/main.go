package main
import "fmt"

func main() {
	var instituteName 		= "NareshLabs"
	const numberOfStudents	= 15
	var studentsPaid		= 10

	fmt.Println("Welcome to", instituteName, "to learn Devops..!!!")
	fmt.Println("Number of students paid", studentsPaid, "from", numberOfStudents)

	fmt.Printf("Welcome to %v to learn Devops..!!!\n", instituteName)
	fmt.Printf("Number of students paid %v from %v \n", studentsPaid, numberOfStudents)

	// Register students for Devops course
	var studentName string
	var amountPaid int

	studentName	= "Anusha"
	amountPaid	= 10000
	fmt.Printf("Student names is: %v, and paid %v \n", studentName, amountPaid)

	// %v for variable calling which defined in Printf
	// %T define varaible Type

	fmt.Printf("instituteName is %T, numberOfStudents is %T and studentsPaid is %T \n", instituteName, numberOfStudents, studentsPaid)

	// Other way to define variables
	var instName string	= "NareshLabs"
	const numberOfStud int  = 15
	var studPaid int		= 10

	fmt.Printf("Welcome to %v to learn Devops..!!! \n", instName)
	fmt.Printf("Number of students paid %v from %v \n", studPaid, numberOfStud)

	// Another way to define variables
	instNameAssign 		:= "NareshLabs"
	numberOfStudAssign	:= 15
	studPaidAssgin 		:= 10

	fmt.Printf("Welcome to %v to learn Devops..!!! \n", instNameAssign)
	fmt.Printf("Number of students paid %v from %v \n", studPaidAssgin, numberOfStudAssign)	
}
