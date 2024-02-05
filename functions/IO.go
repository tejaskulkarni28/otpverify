package functions

import "fmt"

func Greetings(){
	fmt.Println("Welcome to GoLang! OTP verification project")
}

func Inputs() (string, string, string){
	var name,phone,message string

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your phone: ")
	fmt.Scanln(&phone)

	if name!="" && phone!=""{
		message="Thank you for the information."
	}else{
		Inputs()
	}
	return name,phone,message
}

func Outputs(name, phone, message string){
	fmt.Println("Your name is " + name + " and your phone number is " + phone)
	fmt.Println(message)
}