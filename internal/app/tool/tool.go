package tool

import (
	"errors"
	"net/http"
)

//go:generate mockgen -destination=../mocks/auth_tester.go -package=mocks github.com/anatollupacescu/auth-test/internal/main/tool AuthTester
type AuthTester interface {
	TestRequiresAuthorization() (bool, error)
	ResponseCode() (int, error)
}

type Endpoint struct {
	url string
}

func (e *Endpoint) TestRequiresAuthorization() (bool, error) {
	response, err := e.ResponseCode()
	if err != nil {
		return false, err
	}
	return response == 401, nil
}

func (e *Endpoint) ResponseCode() (int, error) {
	response, err := http.Get(e.url)
	if err != nil {
		return -1, err
	}
	return response.StatusCode, nil
}

func NewEndpoint(s string) (AuthTester, error) {
	if len(s) > 0 {
		return &Endpoint{s}, nil
	}
	return nil, errors.New("endpoint url should not be empty")
}
