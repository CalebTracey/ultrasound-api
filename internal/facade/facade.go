package facade

import (
	"github.com/CalebTracey/ultrasound-api/internal/services/auth"
	"github.com/CalebTracey/ultrasound-services/aws"
	"github.com/CalebTracey/ultrasound-services/jwt"
	"github.com/CalebTracey/ultrasound-services/models/res"
	"github.com/CalebTracey/ultrasound-services/mongo"
	"net/http"
	"strconv"
)

type UltrasoundFacadeI interface {
	Authentication(username string) (*res.UserAuthResponse, int)
}

type Receiver struct {
	awsService   aws.ServiceI
	mongoService mongo.ServiceI
	jwtService   jwt.ServiceI
}

func NewUltrasoundService() (Receiver, error) {
	awsSvc := aws.InitializeAwsService()
	mongoSvc := mongo.InitializeMongoService()
	jwtSvc := jwt.InitializeJwtService()
	return Receiver{
		awsService:   awsSvc,
		mongoService: mongoSvc,
		jwtService:   jwtSvc,
	}, nil
}

func (r Receiver) Authentication(username string) (*res.UserAuthResponse, int) {
	vErr := auth.Validate(username)
	if vErr != nil {
		errLogs := errorLogs(http.StatusBadRequest, []error{vErr}, "Validation error")
		return authErrRes(errLogs, 0)
	}
	return &res.UserAuthResponse{
		Success:  true,
		Username: username,
	}, http.StatusOK
}

func errorLogs(code int, errs []error, rc string) (errLogs []res.ErrorLog) {
	for _, e := range errs {
		errLogs = append(errLogs, res.ErrorLog{
			RootCause:  rc,
			StatusCode: strconv.Itoa(code),
			Trace:      e.Error(),
		})
	}
	return errLogs
}
