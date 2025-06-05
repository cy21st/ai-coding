package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ai-go/database"
	"ai-go/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/login", Login)
	return r
}

func setupTestUser() error {
	// Create a test user
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.AdminUser{
		Username: "testuser",
		Password: string(hashedPassword),
		Role:     "admin",
	}

	return database.DB.Create(&user).Error
}

func TestLogin(t *testing.T) {
	// Setup test database
	database.InitDB()
	setupTestUser()

	tests := []struct {
		name       string
		payload    LoginRequest
		wantStatus int
		wantCode   int
	}{
		{
			name: "successful login",
			payload: LoginRequest{
				Username: "testuser",
				Password: "test123",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "wrong password",
			payload: LoginRequest{
				Username: "testuser",
				Password: "wrongpass",
			},
			wantStatus: http.StatusUnauthorized,
			wantCode:   401,
		},
		{
			name: "user not found",
			payload: LoginRequest{
				Username: "nonexistent",
				Password: "test123",
			},
			wantStatus: http.StatusUnauthorized,
			wantCode:   401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupTestRouter()

			// Create request
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assert response
			assert.Equal(t, tt.wantStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCode, int(response["code"].(float64)))

			// If login successful, check token
			if tt.wantStatus == http.StatusOK {
				data := response["data"].(map[string]interface{})
				assert.NotEmpty(t, data["token"])
				user := data["user"].(map[string]interface{})
				assert.Equal(t, "testuser", user["username"])
				assert.Equal(t, "admin", user["role"])
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	router := setupTestRouter()
	router.POST("/api/users", CreateUser)

	tests := []struct {
		name       string
		payload    UserCreateRequest
		wantStatus int
		wantCode   int
	}{
		{
			name: "create user success",
			payload: UserCreateRequest{
				Username: "newuser",
				Password: "password123",
				Role:     "editor",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "missing required fields",
			payload: UserCreateRequest{
				Username: "newuser",
				// Missing password and role
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.wantStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantCode, int(response["code"].(float64)))
		})
	}
}
