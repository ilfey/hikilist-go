package resx

var (
	ResponseOK      = JSON(200, J{"detail": "OK"})
	ResponseCreated = JSON(201, J{"detail": "Created"})

	ResponseBadRequest   = JSON(400, J{"detail": "Bad request"})
	ResponseUnauthorized = JSON(401, J{"detail": "Unauthorized"})
	ResponseForbidden    = JSON(403, J{"detail": "Forbidden"})
	ResponseNotFound     = JSON(404, J{"detail": "Not found"})

	ResponseInternalServerError = JSON(500, J{"detail": "Internal server error"})
)
