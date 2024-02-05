package main
// Download the helper library from https://www.twilio.com/docs/go/install
import ( 
    "tejaskulkarni28/functions"
)

func main(){
	functions.Greetings()
	var name, phone, message = functions.Inputs()
	functions.Outputs(name, phone, message)

	fmt.Println(verification)
}
