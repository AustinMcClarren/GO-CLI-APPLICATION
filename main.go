package main

//IMPORTING PACKAGES
import (
	"fmt"
	"strings"
)

func main (){
conferenceName := "Go Conference" 

const conferenceTickets int  = 50

var remaningTickets uint = 50

bookings := []string{} //EMPTY SLICE ASSIGMENT

	fmt.Printf("conferenceTickets is %T, remaningTickets is %T, conference is %T\n", conferenceTickets,remaningTickets,conferenceName )

	fmt.Printf("Welcome to %v booking \n", conferenceName)

	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remaningTickets)

	fmt.Println("Get your tickets here to attend")

	for {
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
		bookings = append(bookings, firstName + " " + lastName) //SLICING 
	
	
		fmt.Printf("Thank you %v %v for boooking %v tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)
		fmt.Printf("%v tickets remaining for %v\n", remaningTickets,conferenceName)
		
		firstNames:= []string{}
		//this loop ends wehn iterated over all elements of the bookings list
		for _, booking :=  range bookings{
			var names = strings.Fields(booking)
			firstNames = append(firstNames,names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n",firstNames) //list of users who bought tickets 
	}

	}

	
 
	
 

