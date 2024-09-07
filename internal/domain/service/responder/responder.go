package responder

import (
	"encoding/json"
	"github.com/ilfey/hikilist-go/internal/domain/errtype"
	errtypeInterface "github.com/ilfey/hikilist-go/internal/domain/errtype/interface"
	loggerInterface "github.com/ilfey/hikilist-go/pkg/logger/interface"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

type DataResponse struct {
	Data any `json:"data"`
}

type ErrorResponse struct {
	Error error `json:"error"`
}

type Responder struct {
	log loggerInterface.Logger
}

func NewResponder(log loggerInterface.Logger) *Responder {
	return &Responder{
		log: log,
	}
}

func (r *Responder) Respond(w io.Writer, dataOrErr any) {
	err, isErr := dataOrErr.(error)
	if isErr {
		r.log.Error(err)

		var publicErr errtypeInterface.PublicError
		if errors.As(err, &publicErr) {
			// handle the case when write is http.ResponseWriter
			if httpWriter, ok := w.(http.ResponseWriter); ok {
				httpWriter.WriteHeader(publicErr.Status())
			}

			if _, err = w.Write(
				r.toBytes(
					&ErrorResponse{publicErr},
				),
			); err != nil {
				r.log.Critical(err)
			}
		} else {
			// handle the case when write is http.ResponseWriter
			if httpWriter, ok := w.(http.ResponseWriter); ok {
				httpWriter.WriteHeader(http.StatusInternalServerError)
			}

			if _, err = w.Write(
				r.toBytes(
					&ErrorResponse{
						errtype.NewInternalServerError(),
					},
				),
			); err != nil {
				r.log.Critical(err)
			}
		}
		return
	}

	// building a new response data
	resp := &DataResponse{dataOrErr}

	// writing a response data
	if _, err = w.Write(r.toBytes(resp)); err != nil {
		r.log.Critical(err)
	}

	// logging a response
	r.logResponse(resp)
}

func (r *Responder) logResponse(resp *DataResponse) {
	r.log.Object(
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
		r.log.Critical(err)
	}
	return bytes
}
