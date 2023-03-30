package main

import (
	"fmt"
	"net/http")
	

	func index(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"hello, world")
	}

	func about(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"this is my about page")
	}



	func main(){
		http.HandleFunc("/", index)
		http.HandleFunc("/about", about)
		fmt.Println("Serving Starting...")
		http.ListenAndServe(":3001",nil)
	}