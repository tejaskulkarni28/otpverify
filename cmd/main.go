package main
import "fmt"
func inputs() (string, string, string){
	var name,phone,message string

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)

	fmt.Println("Enter your phone: ")
	fmt.Scanln(&phone)

	if name!="" && phone!=""{
		message="Thank you for the information."
	}else{
		inputs()
	}
	return name,phone,message
}

func main(){
	fmt.Println("Welcome to golang! OTP verification project")

	var name, phone, message = inputs()
	fmt.Println("Your name is " + name + " and your phone number is " + phone)
	fmt.Println(message)
}
