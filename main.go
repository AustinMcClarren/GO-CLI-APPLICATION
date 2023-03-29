package main

import "fmt"

func main (){
conferenceName := "Go Conference" 

const conferenceTickets int  = 50

var remaningTickets uint = 50

var bookings [50]string

	fmt.Printf("conferenceTickets is %T, remaningTickets is %T, conference is %T\n", conferenceTickets,remaningTickets,conferenceName )

	fmt.Printf("Welcome to %v booking \n", conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remaningTickets)

	fmt.Println("Get your tickets here to attend")

	

	//user data in var
  	var firstName string

	var lastName string

	var email string

	var userTickets uint

  	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	//ask user for last name
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	//ask user for email
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	//asking user how many tickets they want
	fmt.Println("How many tickets would you like?: ")
	fmt.Scan(&userTickets)
	
	remaningTickets = remaningTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array: %v\n",bookings)
	fmt.Printf("The first value: %v\n",bookings[0])
	fmt.Printf("Array type: %T\n",bookings)
	fmt.Printf("Array length: %v\n",len(bookings))

	fmt.Printf("Thank you %v %v for boooking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
	fmt.Printf("%v tickets remaining for %v\n", remaningTickets,conferenceName)
}
 

