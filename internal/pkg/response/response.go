package response

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ListResponseWithoutGenerics struct {
	Data []any `json:"data"`
}

type ListResponse[T any] struct {
	Data []T `json:"data"`
}
