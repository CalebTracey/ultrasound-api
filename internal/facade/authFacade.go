package facade

import (
	"github.com/CalebTracey/ultrasound-services/models/res"
	"strconv"
)

func authErrRes(errLogs []res.ErrorLog, count int) (*res.UserAuthResponse, int) {
	errLog := errLogs[0]
	status, _ := strconv.Atoi(errLog.StatusCode)
	return &res.UserAuthResponse{
		Success: false,
		Message: res.Message{
			ErrorLog: errLogs,
			Status:   errLog.StatusCode,
			Count:    count,
		},
	}, status
}
