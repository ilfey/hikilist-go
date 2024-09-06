package errtype

//import (
//	"github.com/ilfey/hikilist-go/pkg/logger"
//	"net/http"
//)
//
//const (
//	accessErrType         = "access"
//	publicAccessErrLevel  = logger.InfoLevel
//	publicAccessErrStatus = http.StatusForbidden
//)
//
//type AccessDeniedError struct{ publicError }
//
//func NewAccessDeniedError(msg string) *AccessDeniedError {
//	baseMsg := "access denied"
//	if msg != "" {
//		msg = baseMsg + ": " + msg
//	} else {
//		msg = baseMsg
//	}
//
//	return &AccessDeniedError{
//		publicError{
//			errored{
//				ErrorDetail: msg,
//				ErrorType:   accessErrType,
//				errorLevel:  publicAccessErrLevel,
//				errorStatus: publicAccessErrStatus,
//			},
//		},
//	}
//}
