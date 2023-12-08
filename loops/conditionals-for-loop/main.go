package main
import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string 	= "Go Congerence"
	const conferenceTickets int = 50
	var remainingTickets uint	= 50
	var bookings []string

	fmt.Printf("Welcome to %v bookings application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	for remainingTickets > 0 && len(bookings) < 50 {
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

		if userTickets <= remainingTickets {
			remainingTickets	=	remainingTickets - userTickets
			bookings = append(bookings, firstName+ " " + lastName)
		
			fmt.Printf("Thank you %v %v for bookings %v tickets. you will receive a confirmation to the email at %v\n", firstName, lastName, email, userTickets)
			fmt.Printf("remaining number of tickets %v\n", remainingTickets)

			firstNames	:= []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Our Conference is booked out. Come to next year")
				break
			}
		} else {
			fmt.Printf("we have only %v these many tickets, so can't be book %v tickets\n", remainingTickets, userTickets)
		}
	}
}
