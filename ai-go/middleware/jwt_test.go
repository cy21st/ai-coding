package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"ai-go/config"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestJWTAuth(t *testing.T) {
	// Setup test configuration
	config.GlobalConfig.JWT.SecretKey = "test-secret-key"

	// Generate a test token
	token, err := GenerateToken(1, "testuser", "admin")
	assert.NoError(t, err)

	tests := []struct {
		name           string
		authHeader     string
		wantStatus     int
		wantUserID     uint64
		wantUsername   string
		wantRole       string
		shouldHaveAuth bool
	}{
		{
			name:           "valid token",
			authHeader:     "Bearer " + token,
			wantStatus:     http.StatusOK,
			wantUserID:     1,
			wantUsername:   "testuser",
			wantRole:       "admin",
			shouldHaveAuth: true,
		},
		{
			name:           "missing token",
			authHeader:     "",
			wantStatus:     http.StatusUnauthorized,
			shouldHaveAuth: false,
		},
		{
			name:           "invalid token format",
			authHeader:     "InvalidFormat",
			wantStatus:     http.StatusUnauthorized,
			shouldHaveAuth: false,
		},
		{
			name:           "invalid token",
			authHeader:     "Bearer invalid.token.here",
			wantStatus:     http.StatusUnauthorized,
			shouldHaveAuth: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := setupTestRouter()

			// Add test endpoint
			router.GET("/test", JWTAuth(), func(c *gin.Context) {
				if tt.shouldHaveAuth {
					userID, _ := c.Get("userID")
					username, _ := c.Get("username")
					role, _ := c.Get("role")

					assert.Equal(t, tt.wantUserID, userID)
					assert.Equal(t, tt.wantUsername, username)
					assert.Equal(t, tt.wantRole, role)
				}
				c.Status(http.StatusOK)
			})

			// Create request
			req, _ := http.NewRequest("GET", "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			// Create response recorder
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			// Assert response
			assert.Equal(t, tt.wantStatus, w.Code)
		})
	}
}

func TestGenerateToken(t *testing.T) {
	// Setup test configuration
	config.GlobalConfig.JWT.SecretKey = "test-secret-key"

	tests := []struct {
		name      string
		userID    uint64
		username  string
		role      string
		wantError bool
	}{
		{
			name:      "valid token generation",
			userID:    1,
			username:  "testuser",
			role:      "admin",
			wantError: false,
		},
		{
			name:      "empty username",
			userID:    1,
			username:  "",
			role:      "admin",
			wantError: false, // Should still work as username is not required for token generation
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := GenerateToken(tt.userID, tt.username, tt.role)

			if tt.wantError {
				assert.Error(t, err)
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
