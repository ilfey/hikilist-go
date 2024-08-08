package responses

type J = map[string]any

// ExtensionFunc позволяет расширить ответ. Возвращает расширенный ответ и статус код
type ExtensionFunc = func(...J) (J, int)

func makeResponse(defaultData J, code int) ExtensionFunc {
	return func(extension ...J) (J, int) {
		if len(extension) > 0 {
			for k, v := range extension[0] {
				defaultData[k] = v
			}
		}

		return defaultData, code
	}
}

var (
	ResponseOK      = makeResponse(J{"detail": "OK"}, 200)
	ResponseCreated = makeResponse(J{"detail": "Created"}, 201)

	ResponseBadRequest       = makeResponse(J{"detail": "Bad request"}, 400)
	ResponseUnauthorized     = makeResponse(J{"detail": "Unauthorized"}, 401)
	ResponseForbidden        = makeResponse(J{"detail": "Forbidden"}, 403)
	ResponseNotFound         = makeResponse(J{"detail": "Not found"}, 404)
	ResponseMethodNotAllowed = makeResponse(J{"detail": "Method not allowed"}, 405)

	ResponseInternalServerError = makeResponse(J{"detail": "Internal server error"}, 500)
)
