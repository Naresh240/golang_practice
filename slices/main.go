package main
import "fmt"

func main() {
	var conferenceName string 	= "Go Congerence"
	const conferenceTickets int = 50
	var remainingTickets uint	= 50

	fmt.Printf("Welcome to %v bookings application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("please enter first name")
	fmt.Scan(&firstName)
	fmt.Println("please enter last name")
	fmt.Scan(&lastName)
	fmt.Println("please enter email name")
	fmt.Scan(&email)
	fmt.Println("please enter number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets	=	remainingTickets - userTickets

	fmt.Printf("Thank you %v %v for bookings %v tickets. you will receive a confirmation to the email at %v\n", firstName, lastName, email, userTickets)
	fmt.Printf("remaining number of tickets %v\n", remainingTickets)

	//	Define and Re-use slice varaible
	// var bookings = []string{}
	// var bookings := []string{}
	var bookings []string
	bookings = append(bookings, firstName+ " " + lastName)

	fmt.Printf("The whole bookings slice %v\n", bookings)
	fmt.Printf("The first value of arragy is %v\n", bookings[0])
	fmt.Printf("Slice Type %T\n", bookings)
	fmt.Printf("Slice Length %v\n", len(bookings))	
}
