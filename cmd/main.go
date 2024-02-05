package main

import ( 
    "tejaskulkarni28/functions"
)

func main(){
	functions.Greetings()
	var name, phone, message = functions.Inputs()
	functions.Outputs(name, phone, message)
}
