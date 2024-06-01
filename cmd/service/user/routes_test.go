package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/HarshitPG/ecommerce_api_go/cmd/types"
	"github.com/gorilla/mux"
)

func TestUserService(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "UserTesing",
			LastName:  "Testing",
			Email:     "123@gmail.com",
			Password:  "t12",
		}

		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, got %d", http.StatusBadGateway, rr.Code)
		}

	})
}

type mockUserStore struct{}


func (m *mockUserStore) CreateUser(types.User) error {
	return nil
}


func (m *mockUserStore) GetuserByID(id int) (*types.User, error) {
	return nil,nil
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}
