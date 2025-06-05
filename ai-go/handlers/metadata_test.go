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
)

// 定义请求结构体
// type RelationCreateRequest struct {
// 	EventID uint64 `json:"event_id" binding:"required"`
// 	AttrID  uint64 `json:"attr_id" binding:"required"`
// }

func setupTestMetadataRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func setupTestMetadata() error {
	// Create test event
	event := models.MetaEvent{
		EventName: "test_event",
		EventDesc: "Test event description",
		IsDeleted: false,
	}
	if err := database.DB.Create(&event).Error; err != nil {
		return err
	}

	// Create test attribute
	attr := models.MetaAttr{
		AttrName:  "test_attr",
		AttrDesc:  "Test attribute description",
		AttrType:  "string",
		IsDeleted: false,
	}
	if err := database.DB.Create(&attr).Error; err != nil {
		return err
	}

	// Create test relation
	relation := models.MetaRelation{
		EventID: event.ID,
		AttrID:  attr.ID,
	}
	return database.DB.Create(&relation).Error
}

func TestCreateEvent(t *testing.T) {
	// Setup test database
	database.InitDB()

	router := setupTestMetadataRouter()
	router.POST("/api/events", CreateEvent)

	tests := []struct {
		name       string
		payload    EventCreateRequest
		wantStatus int
		wantCode   int
	}{
		{
			name: "create event success",
			payload: EventCreateRequest{
				EventName: "new_event",
				EventDesc: "New event description",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "missing required fields",
			payload: EventCreateRequest{
				EventName: "",
				EventDesc: "New event description",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/api/events", bytes.NewBuffer(jsonData))
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

func TestCreateAttribute(t *testing.T) {
	router := setupTestMetadataRouter()
	router.POST("/api/attributes", CreateAttribute)

	tests := []struct {
		name       string
		payload    AttributeCreateRequest
		wantStatus int
		wantCode   int
	}{
		{
			name: "create attribute success",
			payload: AttributeCreateRequest{
				AttrName: "new_attr",
				AttrDesc: "New attribute description",
				AttrType: "string",
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "missing required fields",
			payload: AttributeCreateRequest{
				AttrName: "",
				AttrDesc: "New attribute description",
				AttrType: "string",
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/api/attributes", bytes.NewBuffer(jsonData))
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

func TestCreateRelation(t *testing.T) {
	// Setup test database and data
	database.InitDB()
	setupTestMetadata()

	router := setupTestMetadataRouter()
	router.POST("/api/relations", CreateRelation)

	tests := []struct {
		name       string
		payload    RelationCreateRequest
		wantStatus int
		wantCode   int
	}{
		{
			name: "create relation success",
			payload: RelationCreateRequest{
				EventID: 1,
				AttrID:  1,
			},
			wantStatus: http.StatusOK,
			wantCode:   200,
		},
		{
			name: "missing required fields",
			payload: RelationCreateRequest{
				EventID: 0,
				AttrID:  1,
			},
			wantStatus: http.StatusBadRequest,
			wantCode:   400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonData, _ := json.Marshal(tt.payload)
			req, _ := http.NewRequest("POST", "/api/relations", bytes.NewBuffer(jsonData))
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

func TestGetEventList(t *testing.T) {
	// Setup test database and data
	database.InitDB()
	setupTestMetadata()

	router := setupTestMetadataRouter()
	router.GET("/api/events", GetEventList)

	t.Run("get event list", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/events", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 200, int(response["code"].(float64)))

		data := response["data"].([]interface{})
		assert.Greater(t, len(data), 0)
	})
}

func TestGetAttributeList(t *testing.T) {
	router := setupTestMetadataRouter()
	router.GET("/api/attributes", GetAttributeList)

	t.Run("get attribute list", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/attributes", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 200, int(response["code"].(float64)))

		data := response["data"].([]interface{})
		assert.Greater(t, len(data), 0)
	})
}

func TestGetRelationList(t *testing.T) {
	router := setupTestMetadataRouter()
	router.GET("/api/relations", GetRelationList)

	t.Run("get relation list", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/relations", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 200, int(response["code"].(float64)))

		data := response["data"].([]interface{})
		assert.Greater(t, len(data), 0)
	})
}
