package common

// A response to use in json response when meta data is needed
type MetaResponse struct {
	Data any `json:"data"`
	Meta any `json:"meta"`
}
