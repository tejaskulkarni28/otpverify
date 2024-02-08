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
	service_sid := os.Getenv("VERIFY_SERVICE_SID")
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	if phone != ""{
		to := "+91" + phone

		clientWithParams := twilio.NewRestClientWithParams(twilio.ClientParams{
			Username: accountSid,
			Password: authToken,
		})

		params := &verification.CreateVerificationParams{}
		params.SetTo(to)
		params.SetChannel("sms")

		res, err := clientWithParams.VerifyV2.CreateVerification(service_sid, params)
		if err != nil{
			fmt.Println(err.Error())
		}else{
			VerifyOTP(*res.Status)
		}
	}else{
		fmt.Println("Phone number not recieved! ")
	}
}
