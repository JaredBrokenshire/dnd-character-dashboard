package tests

import (
	"dnd-character-dashboard/api/requests"
	"dnd-character-dashboard/services"
	"dnd-character-dashboard/tests/helpers"
	"dnd-character-dashboard/tests/mocks"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
)

func TestAuth_Register(t *testing.T) {
	ts.ClearTable("users")
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/register",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't register without the required fields",
			Request:     request,
			RequestBody: requests.RegisterRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username is a required field",
					"Email is a required field",
					"Password is a required field",
					"ConfirmPassword is a required field",
				},
			},
		},
		{
			Name:    "Can't register with an existing username",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        ts.TestUser.Username,
				Email:           "test-email@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Username is already taken",
			},
		},
		{
			Name:    "Can't register with an existing email",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           ts.TestUser.Email,
				Password:        "password",
				ConfirmPassword: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Email is already registered",
			},
		},
		{
			Name:    "Can't register with a password that doesn't match the confirm password",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           "test-email@example.com",
				Password:        "password",
				ConfirmPassword: "password2",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyPart:   "Passwords do not match",
			},
		},
		{
			Name: "Can't register if there is an error from the auth service",
			Setup: func(test *helpers.TestCase) {
				mockAuthService := mocks.NewMockAuthService()
				mockAuthService.RegisterError = errors.New("test error")

				ts.S.Dependencies.SetAuthService(mockAuthService)
			},
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           "test-email@example.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error registering",
			},
			Teardown: func() {
				ts.S.Dependencies.SetAuthService(
					services.NewAuthService(
						ts.S.DB,
						ts.S.Repos.User,
						ts.S.Repos.Role,
						os.Getenv("JWT_SECRET"),
					),
				)
			},
		},
		{
			Name:    "Can register with valid details",
			Request: request,
			RequestBody: requests.RegisterRequest{
				Username:        "test-username",
				Email:           "test-email@example.com",
				Password:        "test-password",
				ConfirmPassword: "test-password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"access_token":"`,
					`"username":"test-username"`,
					`"email":"test-email@example.com"`,
				},
				BodyPartsMissing: []string{
					"test-password",
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}

func TestAuth_Login(t *testing.T) {
	ts.ClearTable("users")
	ts.SetupDefaultUsers()

	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/auth/login",
	}

	cases := []helpers.TestCase{
		{
			Name:        "Can't login without the required fields",
			Request:     request,
			RequestBody: requests.LoginRequest{},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusBadRequest,
				BodyParts: []string{
					"Required fields are empty or not valid",
					"Username is a required field",
					"Password is a required field",
				},
			},
		},
		{
			Name:    "Can't login with username that doesn't exist",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: "Invalid Username",
				Password: "Invalid Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Invalid credentials",
			},
		},
		{
			Name:    "Can't login with incorrect password",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: "Incorrect Password",
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusUnauthorized,
				BodyPart:   "Invalid credentials",
			},
		},
		{
			Name: "Can't login if there is an error from the auth service",
			Setup: func(test *helpers.TestCase) {
				mockAuthService := mocks.NewMockAuthService()
				mockAuthService.LoginError = errors.New("test error")

				ts.S.Dependencies.SetAuthService(mockAuthService)
			},
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: helpers.AdminPassword,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusInternalServerError,
				BodyPart:   "There was an error logging in",
			},
			Teardown: func() {
				ts.S.Dependencies.SetAuthService(
					services.NewAuthService(
						ts.S.DB,
						ts.S.Repos.User,
						ts.S.Repos.Role,
						os.Getenv("JWT_SECRET"),
					),
				)
			},
		},
		{
			Name:    "Can login with correct credentials",
			Request: request,
			RequestBody: requests.LoginRequest{
				Username: helpers.AdminUsername,
				Password: helpers.AdminPassword,
			},
			Expected: helpers.ExpectedResponse{
				StatusCode: http.StatusOK,
				BodyParts: []string{
					`"access_token":"`,
					fmt.Sprintf(`"username":"%v"`, ts.AdminUser.Username),
				},
				BodyPartsMissing: []string{
					helpers.AdminPassword,
				},
			},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			RunAuthorisedTestCase(t, test)
		})
	}
}
