package main

import "fmt"

func main (){
var conferenceName string = "Go Conference" 
const conferenceTickets int  = 50
var remaningTickets uint = 50

	fmt.Printf("conferenceTickets is %T, remaningTickets is %T, conference is %T\n", conferenceTickets,remaningTickets,conferenceName )


	fmt.Printf("Welcome to %v booking \n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remaningTickets)
	fmt.Println("Get your tickets here to attend")

  	var userName string
	var userTickets int
  	//ask user for their name

	userName = "Austin"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
 

