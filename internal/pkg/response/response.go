package response

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ListResponse[T any] struct {
	Data []T `json:"data"`
}
