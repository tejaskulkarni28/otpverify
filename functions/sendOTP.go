package functions
import (
	"fmt"
	"github.com/twilio/twilio-go"
	verification "github.com/twilio/twilio-go/rest/verify/v2"
	"github.com/joho/godotenv"
    "os"
)

func SendOTP(phone string){

	err := godotenv.Load("auth/.env")
	if err != nil{
		fmt.Println("Error loading env file", err)
		return
	}
	// accesing env variables
	service_sid := os.Getenv("VERIFY_SERVICE_SID")

	// +16562230173
	if phone != ""{
		to := "+1" + phone
		fmt.Println(to)

		// created a new twilio rest client 
		client := twilio.NewRestClient()

		// creating parameters and initializing 
		params := &verification.CreateVerificationParams{}

		// setting phone number
		params.SetTo(to)
		params.SetChannel("sms")

		res, err:= client.VerifyV2.CreateVerification(service_sid, params)
		if err != nil{
			fmt.Println(err.Error())
		}else{
			fmt.Println(res)
			// if res.Sid != nil {
			// 	fmt.Println(*res.Sid)
			// } else {
			// 	fmt.Println(res)
			// }
		}
	}else{
		fmt.Println("Phone number not recieved! ")
	}
}