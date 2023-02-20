package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

// main shows how to create a context and pass it to a function
func main() {
	ctx := context.Background()
	result, err := logic(ctx, "a string")
	fmt.Println(result, err)
}

// logic shows the parameters for functions that pass or use the context
func logic(ctx context.Context, info string) (string, error) {
	// do some interesting stuff here
	return "", nil
}

// Middleware shows what middleware wrappers look like when the place values into a context
func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		// wrap the context with stuff -- we'll see how soon!
		req = req.WithContext(ctx)
		handler.ServeHTTP(rw, req)
	})
}

// handler shows how to extract a context from an *http.Request and pass it to a function
func handler(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := req.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	data := req.FormValue("data")
	result, err := logic(ctx, data)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}
	rw.Write([]byte(result))
}

type ServiceCaller struct {
	client *http.Client
}

// callAnotherService shows how to add a context to an *http.Request
func (sc ServiceCaller) callAnotherService(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://example.com?data="+data, nil)
	if err != nil {
		return "", err
	}
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d",
			resp.StatusCode)
	}
	// do the rest of the stuff to process the response
	id, err := processResponse(resp.Body)
	return id, err
}

// processResponse is a placeholder function for processing the body of an *http.Response
func processResponse(body io.ReadCloser) (string, error) {
	return "", nil
}
