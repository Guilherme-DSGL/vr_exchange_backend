package rest

import "net/http"

type HttpClient interface {
	// DO represents the GET http request
	Do(req *http.Request) (*http.Response, error)
}
