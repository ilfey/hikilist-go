package responder

import (
	"encoding/json"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	errtypeInterface "github.com/ilfey/hikilist-go/internal/domain/errtype/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"io"
	"net/http"
	"time"
)

type DataResponse struct {
	Data any `json:"data"`
}

func NewDataResponse(data any) DataResponse {
	return DataResponse{Data: data}
}

type ErrorResponse struct {
	Error error `json:"error"`
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{Error: err}
}

type Responder struct {
	logger loggerInterface.Logger
}

func NewResponder(logger loggerInterface.Logger) *Responder {
	return &Responder{
		logger: logger,
	}
}

func (r *Responder) Respond(w io.Writer, dataOrErr any) {
	err, isErr := dataOrErr.(error)
	if isErr {
		r.logger.Log(err)

		publicErr, isPublicErr := err.(errtypeInterface.PublicError)
		if isPublicErr {
			// handle the case when write is http.ResponseWriter
			if httpWriter, ok := w.(http.ResponseWriter); ok {
				httpWriter.WriteHeader(publicErr.Status())
			}

			if _, err = w.Write(
				r.toBytes(
					NewErrorResponse(publicErr),
				),
			); err != nil {
				r.logger.Critical(err)
			}
		} else {
			// handle the case when write is http.ResponseWriter
			if httpWriter, ok := w.(http.ResponseWriter); ok {
				httpWriter.WriteHeader(http.StatusInternalServerError)
			}

			if _, err = w.Write(
				r.toBytes(
					NewErrorResponse(
						errtype.NewInternalServerError(),
					),
				),
			); err != nil {
				r.logger.Critical(err)
			}
		}
		return
	}

	// building a new response data
	resp := NewDataResponse(dataOrErr)

	// writing a response data
	if _, err = w.Write(r.toBytes(resp)); err != nil {
		r.logger.Critical(err)
	}

	// logging a response
	r.logResponse(resp)
}

func (r *Responder) logResponse(resp DataResponse) {
	r.logger.LogData(
		&LoggableData{
			Date:         time.Now(),
			Type:         LogType,
			DataResponse: resp,
		},
	)
}

func (r *Responder) toBytes(resp any) []byte {
	bytes, err := json.Marshal(resp)
	if err != nil {
		r.logger.Critical(err)
	}
	return bytes
}
