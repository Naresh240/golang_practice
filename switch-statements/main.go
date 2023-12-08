package main
import (
	"fmt"
)

func main() {
	fmt.Printf("Please provide city name...!!!!\n")

	var city string
	fmt.Scan(&city)

	switch city {
		case "Bangalore":
			fmt.Println("I am working at", city, "and I Like to be settle in", city)
		case "Hyderabad":
			fmt.Println("I am working at", city, "and I Like to be settle in", city)
		case "Amaravathi":
			fmt.Println("I am working at", city, "and I Like to be settle in", city)
		case "Mumbai":
			fmt.Println("I am working at", city, "and I Like to be settle in", city)
		case "Delhi":
			fmt.Println("I am working at", city, "and I Like to be settle in", city)
		case "Chennai":
			fmt.Println("I am working at", city, "and I Like to be settle in", city)
		default:
			fmt.Printf("Not a valid city selected\n")
	}
}
