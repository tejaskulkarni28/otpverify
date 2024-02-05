package main
// Download the helper library from https://www.twilio.com/docs/go/install
import ( 
    "tejaskulkarni28/functions"
)

func main(){

	// 1. greetings
	functions.Greetings()

	// 2. input 
	var name, phone, message = functions.Inputs()

	// 3. output
	functions.Outputs(name, phone, message)

	// passing output values to sendOTP.go
	functions.SendOTP(phone)


}
