package endpoints_test

import (
	"encoding/json"
	"github.com/lambovg/go-request-compose/examples/endpoints"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/lambovg/go-request-compose/examples/dataprovider"
)

func TestListUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	expectedUsers := []dataprovider.User{
		{ID: 1, Name: "John", Email: "john@localhost"},
		{ID: 2, Name: "Doe ", Email: "doe@localhost"},
	}

	mock := new(dataprovider.UsersMock)
	mock.MockedListUsers = func() ([]dataprovider.User, error) {
		return expectedUsers, nil
	}

	endpoints.ListUsers(mock).ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	var gotUsers []dataprovider.User
	if err := json.NewDecoder(res.Body).Decode(&gotUsers); err != nil {
		t.Errorf("decode response body: %v", err)
	}

	if !reflect.DeepEqual(expectedUsers, gotUsers) {
		t.Errorf("got users { %+v } did not match expected users { %+v }", gotUsers, expectedUsers)
	}
}
