package routes

import (
	"github.com/CalebTracey/ultrasound-services/models/res"
	"os"
	"strconv"
)

func SetAuthenticationResponse(res *res.UserAuthResponse) *res.UserAuthResponse {
	hn, _ := os.Hostname()
	status, _ := strconv.Atoi(res.Message.Status)
	res.Message.Count = 1
	res.Message.HostName = hn
	res.Message.Status = strconv.Itoa(status)
	return res
}
