package responderInterface

import "io"

type Responder interface {
	Respond(w io.Writer, dataOrErr any)
}
