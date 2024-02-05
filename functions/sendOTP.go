package functions
import (
	"fmt"
)

func SendOTP(phone string){
	if phone != ""{
		fmt.Println(phone)
	}else{
		fmt.Println("Phone number not recieved! ")
	}
}