package auth

import (
	"ColdStart/internal/pkg/log"
	"ColdStart/internal/pkg/random"
	"ColdStart/models"
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Auth struct {
	config *viper.Viper
	logger *log.Logger
	db     *gorm.DB
}

func (handler *Auth) createJWTToken(userID string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    userID,
		"ExpiresAt": time.Now().Add(72 * time.Hour).Unix(),
	})
	return claims.SignedString([]byte(handler.config.GetString("security.jwt.secret")))
}

func (handler *Auth) IsAuthenticated(r *http.Request) *models.User {
	// Retrieve JWT cookie
	cookie, err := r.Cookie("jwt")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			handler.logger.Debug("No JWT cookie found")
		}
		return nil
	}

	// Parse JWT token
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(handler.config.GetString("security.jwt.secret")), nil
	})
	if err != nil || !token.Valid {
		handler.logger.Debug("Invalid JWT token")
		return nil
	}

	// Extract claims and ensure "Issuer" exists
	claims, ok := token.Claims.(*jwt.MapClaims)
	if !ok || (*claims)["Issuer"] == nil {
		handler.logger.Debug("Invalid JWT claims")
		return nil
	}

	// Query the database for the user based on the Issuer (user ID)
	var user models.User
	if err := handler.db.Where("id = ?", (*claims)["Issuer"]).First(&user).Error; err != nil {
		handler.logger.Debug("User not found")
		return nil
	}

	handler.logger.Debug("User authenticated", zap.String("UUID", user.Id))
	return &user
}

// Authenticate godoc
// @Summary Authenticate user
// @Description Authenticate the user by checking the JWT token in the request cookies. If authenticated, return user data.
// @Tags Accounts
// @Accept json
// @Produce json
// @Success 200 {object} models.User "Authenticated user data"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /accounts/user/ [get]
func (handler *Auth) Authenticate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check if the user is authenticated
	user := handler.IsAuthenticated(r)
	if user == nil || user.Id == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		handler.logger.Debug("Authentication failed. User not authorized.")
		return
	}

	// Respond with the authenticated user data
	if err := json.NewEncoder(w).Encode(user); err != nil {
		handler.logger.Fatal("Error encoding user data", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	handler.logger.Debug("User authenticated successfully.")
}

// Login godoc
// @Summary User login
// @Description Logs in the user by verifying their email and password. Returns a JWT token and user data upon successful login.
// @Tags Accounts
// @Accept json
// @Produce json
//
// @Param email    body string true "User's email"    example("example@gmail.com")
// @Param password body string true "User's password" example("Pa$$w0rd")
//
// @Success 200 {object} models.User "Logged-in user data"
// @Failure 400 {object} map[string]string "Invalid request data"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /accounts/login [post]
func (handler *Auth) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode the user
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.logger.Debug("Login failed. Could not decode body.")
		http.Error(w, "Invalid request data.", http.StatusBadRequest)
		return
	}

	// Validate input data
	if user.Email == "" || user.Password == "" {
		handler.logger.Debug("Login failed. Missing email or password.")
		http.Error(w, "Email and password are required.", http.StatusBadRequest)
		return
	}

	// Query the database for user by email
	var oldUser models.User
	handler.db.Where("email = ?", user.Email).First(&oldUser)
	if oldUser.Id == "" {
		handler.logger.Debug("Login failed. User not found.")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Check if password matches
	if !matchPassword(oldUser.Password, user.Password, oldUser.Salt) {
		handler.logger.Debug("Login failed. Incorrect password.")
		http.Error(w, "Incorrect password.", http.StatusUnauthorized)
		return
	}

	// Creating new JWT token
	token, err := handler.createJWTToken(oldUser.Id)
	if err != nil {
		handler.logger.Fatal("Error generating JWT token", zap.Error(err))
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HttpOnly: true,
	})

	// Return the user info
	handler.logger.Debug("Login succeeded.")
	if err := json.NewEncoder(w).Encode(oldUser); err != nil {
		handler.logger.Info("Unexpected error occurred while encoding user", zap.Error(err))
	}
}

// Logout godoc
// @Summary User logout
// @Description Logs out the user by invalidating the JWT token stored in the cookie.
// @Tags Accounts
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Logout successful message"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /accounts/logout [post]
func (handler *Auth) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Invalidate the JWT cookie by setting an expired cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	// Logout action
	handler.logger.Debug("User logged out successfully.")

	// Send a success response
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"message": "Logged out"}`))
	if err != nil {
		handler.logger.Debug("failed to write response", zap.Error(err))
		return
	}
}

// Register godoc
// @Summary User registration
// @Description Registers a new user. If the login is empty, returns a temporary UUID. If successful, returns the created user's information, including a UUID, login, email, and other details.
// @Tags Accounts
// @Accept json
// @Produce json
//
// @Param login    body string true "User's login"    example("Example")
// @Param email body string true "User's email" example("example@gmail.com")
// @Param password body string true "User's password" example("Pa$$w0rd")
// @Param name body string false "User's name" example("Fedor")
// @Param surname body string false "User's surname" example("Triphosphate")
// @Param region body string false "User's region" example("CFD")
// @Param city body string false "User's city" example("Moscow")
//
// @Success 200 {object} models.User "Registered user information"
// @Failure 400 {object} map[string]string "Invalid request data or registration failed"
// @Router /accounts/register [post]
func (handler *Auth) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Decode the user
	var user struct {
		Login    string `json:"login"`
		Email    string `json:"email"`
		Password string `json:"password"`

		Name    string `json:"name"`
		Surname string `json:"surname"`
		Region  string `json:"region"`
		City    string `json:"city"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		handler.logger.Debug("Registration failed. Could not decode body.")
		http.Error(w, "Invalid request data.", http.StatusBadRequest)
		return
	}

	// Assuming user is temporary we return a UUID for frontend identification
	if user.Login == "" {
		handler.logger.Debug("User login is empty: proceeding to temporary UUID")
		tempID := uuid.New().String()
		newUser := models.User{
			Id:       tempID,
			Login:    tempID,
			Email:    tempID,
			Salt:     random.GenerateSalt(handler.config.GetInt("security.auth.salt_size")),
			Password: tempID,
		}

		// Creating new user in DB
		if err := handler.db.Create(&newUser).Error; err != nil {
			handler.logger.Debug("Temp authorisation failed", zap.Error(err))
			http.Error(w, "Temp authorisation failed.", http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(newUser); err != nil {
			handler.logger.Fatal("Unexpected error occurred", zap.Error(err))
		}
		return
	}

	// Generate a new UUID for the user
	newUser := models.User{
		Id:    uuid.New().String(),
		Login: user.Login,
		Email: user.Email,
		Salt:  random.GenerateSalt(handler.config.GetInt("security.auth.salt_size")),
	}

	newUser.Password = hashPassword(user.Password, newUser.Salt)

	// Creating new user in DB
	if err := handler.db.Create(&newUser).Error; err != nil {
		handler.logger.Debug("Registration failed", zap.Error(err))
		http.Error(w, "Registration failed.", http.StatusBadRequest)
		return
	}

	handler.logger.Debug("User registered successfully.")

	// Return the user info
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		handler.logger.Fatal("Unexpected error occurred", zap.Error(err))
	}
}

func New(log *log.Logger, cfg *viper.Viper, db *gorm.DB) *Auth {
	log.Info("Initializing auth handler")
	defer func() { log.Info("Auth handler initialized") }()
	return &Auth{
		config: cfg,
		logger: log,
		db:     db,
	}
}
