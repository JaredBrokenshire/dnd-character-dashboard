package handlers

import (
	"dnd-character-dashboard/api"
	"dnd-character-dashboard/api/requests"
	"dnd-character-dashboard/api/responses"
	"dnd-character-dashboard/pkg/logging"
	"dnd-character-dashboard/services"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strings"
)

type AuthHandler struct {
	server *api.Server
}

func NewAuthHandler(server *api.Server) *AuthHandler {
	return &AuthHandler{server: server}
}

// Register godoc
// @Summary New User Register
// @Description Register a new user
// @ID auth-register
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Param params body requests.RegisterRequest true "New User's credentials"
// @Success 200 {object} responses.AuthResponse
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /auth/register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	request := new(requests.RegisterRequest)
	if err := c.Bind(request); err != nil {
		logging.Logger.LogError().Msg("auth/register bind error: " + err.Error())
		return err
	}
	if err := c.Validate(request); err != nil {
		logging.Logger.LogError().Msg("auth/register request validation error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	dto := services.RegisterRequestDTO{
		Username:        request.Username,
		Email:           request.Email,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}

	authResponse, refreshToken, err := h.server.Dependencies.GetAuthService().Register(dto)
	switch {
	case errors.Is(err, services.ErrUsernameTaken):
		logging.Logger.LogWarn().Msg("auth/register duplicate username")
		return responses.ErrorResponse(c, http.StatusBadRequest, "Username is already taken")
	case errors.Is(err, services.ErrEmailTaken):
		logging.Logger.LogWarn().Msg("auth/register duplicate email")
		return responses.ErrorResponse(c, http.StatusBadRequest, "Email is already registered")
	case errors.Is(err, services.ErrPasswordNotMatch):
		logging.Logger.LogWarn().Msg("auth/register password mismatch")
		return responses.ErrorResponse(c, http.StatusBadRequest, "Passwords do not match")
	case err != nil:
		logging.Logger.LogError().Msg("auth/register error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error registering")
	}

	h.setRefreshCookie(c, refreshToken)

	res := responses.NewAuthResponse(authResponse.User, authResponse.AccessToken)
	return responses.Response(c, http.StatusOK, res)
}

// Login godoc
// @Summary User Login
// @Description Log a user in
// @ID auth-login
// @Tags Auth Actions
// @Accept json
// @Produce json
// @Param params body requests.LoginRequest true "User's credentials"
// @Success 200 {object} responses.AuthResponse
// @Failure 400 {object} responses.Error
// @Failure 401 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /auth/login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	request := new(requests.LoginRequest)
	if err := c.Bind(request); err != nil {
		logging.Logger.LogError().Msg("auth/login bind error: " + err.Error())
		return err
	}
	if err := c.Validate(request); err != nil {
		logging.Logger.LogError().Msg("auth/login request validation error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty or not valid: "+err.Error())
	}

	dto := services.LoginRequestDTO{
		Username: request.Username,
		Password: request.Password,
	}

	authResponse, refreshToken, err := h.server.Dependencies.GetAuthService().Login(dto)
	switch {
	case errors.Is(err, services.ErrInvalidCredentials):
		logging.Logger.LogWarn().Msg("auth/login invalid credentials")
		return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid credentials")
	case err != nil:
		logging.Logger.LogError().Msg("auth/login error: " + err.Error())
		return responses.ErrorResponse(c, http.StatusInternalServerError, "There was an error logging in")
	}

	h.setRefreshCookie(c, refreshToken)

	res := responses.NewAuthResponse(authResponse.User, authResponse.AccessToken)
	return responses.Response(c, http.StatusOK, res)
}

func (h *AuthHandler) setRefreshCookie(c echo.Context, token string) {
	// Use Secure: true in production (requires HTTPS). Controlled by ENVIRONMENT env var.
	secureCookie := strings.EqualFold(os.Getenv("ENVIRONMENT"), "production")

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   secureCookie,
		MaxAge:   7 * 24 * 60 * 60,
	}

	c.SetCookie(cookie)
}
