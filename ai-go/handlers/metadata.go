package handlers

import (
	"slices"
	"strconv"
	"strings"
	"time"

	"ai-go/database"
	"ai-go/models"
	"ai-go/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventCreateRequest struct {
	EventName string `json:"event_name" binding:"required"`
	EventDesc string `json:"event_desc"`
}

type AttributeUpdateRequest struct {
	AttrName string  `json:"attr_name" binding:"required"`
	AttrType *string `json:"attr_type"`
	AttrDesc *string `json:"attr_desc"`
}

type AttributeCreateRequest struct {
	AttrName string `json:"attr_name" binding:"required"`
	AttrType string `json:"attr_type" binding:"required"`
	AttrDesc string `json:"attr_desc"`
}

type RelationCreateRequest struct {
	EventID uint64 `json:"event_id" binding:"required"`
	AttrID  uint64 `json:"attr_id" binding:"required"`
}

// RelationDetail 包含关联关系的详细信息
type RelationDetail struct {
	ID         uint64    `json:"id"`
	EventID    uint64    `json:"event_id"`
	AttrID     uint64    `json:"attr_id"`
	EventName  string    `json:"event_name"`
	AttrName   string    `json:"attr_name"`
	IsDeleted  bool      `json:"is_deleted"`
	CreatedAt  time.Time `json:"-"`
	CreatedStr string    `json:"created_at"`
}

// AttributeType 属性类型映射
type AttributeType struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// ValidAttributeTypes 预定义的有效属性类型
var ValidAttributeTypes = []string{"string", "number", "bool", "time"}

// Event handlers
func CreateEvent(c *gin.Context) {
	var req EventCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request parameters", err)
		return
	}

	event := models.MetaEvent{
		EventName: req.EventName,
		EventDesc: req.EventDesc,
	}

	if err := database.DB.Create(&event).Error; err != nil {
		utils.InternalError(c, "Failed to create event", err)
		return
	}

	utils.Success(c, event)
}

func UpdateEvent(c *gin.Context) {
	eventID := c.Param("id")
	var req EventCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request parameters", err)
		return
	}

	var event models.MetaEvent
	if err := database.DB.First(&event, eventID).Error; err != nil {
		utils.NotFound(c, "Event not found", err)
		return
	}

	event.EventName = req.EventName
	event.EventDesc = req.EventDesc

	if err := database.DB.Save(&event).Error; err != nil {
		utils.InternalError(c, "Failed to update event", err)
		return
	}

	utils.Success(c, event)
}

func DeleteEvent(c *gin.Context) {
	eventID := c.Param("id")

	// 使用事务同时软删除事件和关联关系
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 软删除关联关系
		if err := tx.Model(&models.MetaRelation{}).
			Where("event_id = ?", eventID).
			Update("is_deleted", true).Error; err != nil {
			return err
		}

		// 软删除事件
		if err := tx.Model(&models.MetaEvent{}).
			Where("id = ?", eventID).
			Update("is_deleted", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		utils.InternalError(c, "Failed to delete event and its relations", err)
		return
	}

	utils.SuccessWithMessage(c, "Event and its relations deleted successfully", nil)
}

func GetEventList(c *gin.Context) {
	// 获取分页参数
	pageNumStr := c.Query("page_num")
	pageSizeStr := c.Query("page_size")

	// 获取搜索关键词
	keyword := c.Query("keyword")

	var events []models.MetaEvent
	var total int64

	// 创建查询构建器
	query := database.DB.Model(&models.MetaEvent{}).Where("is_deleted = ?", false)

	// 如果有搜索关键词，添加搜索条件
	if keyword != "" {
		query = query.Where("event_name LIKE ? OR event_desc LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总记录数
	query.Count(&total)

	// 如果没有传分页参数，返回所有数据
	if pageNumStr == "" || pageSizeStr == "" {
		err := query.Order("id DESC").Find(&events).Error
		if err != nil {
			utils.InternalError(c, "Failed to get event list", err)
			return
		}

		utils.Success(c, gin.H{
			"events": events,
			"pagination": gin.H{
				"page_num":   1,
				"page_size":  total,
				"total":      total,
				"total_page": 1,
			},
		})
		return
	}

	// 有分页参数时进行分页查询
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 添加排序并使用分页获取数据
	err := utils.Paginate(pageNum, pageSize, query.Order("id DESC")).Find(&events).Error
	if err != nil {
		utils.InternalError(c, "Failed to get event list", err)
		return
	}

	utils.Success(c, gin.H{
		"events": events,
		"pagination": gin.H{
			"page_num":   pageNum,
			"page_size":  pageSize,
			"total":      total,
			"total_page": utils.TotalPage(total, pageSize),
		},
	})
}

func GetEventInfo(c *gin.Context) {
	eventID := c.Param("id")
	var event models.MetaEvent
	if err := database.DB.Preload("Attributes", func(db *gorm.DB) *gorm.DB {
		return db.Order("id DESC")
	}).First(&event, eventID).Error; err != nil {
		utils.NotFound(c, "Event not found", err)
		return
	}

	utils.Success(c, event)
}

// Attribute handlers
func CreateAttribute(c *gin.Context) {
	var req AttributeCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request", err)
		return
	}

	// 验证属性类型是否有效
	if !isValidAttributeType(req.AttrType) {
		utils.BadRequest(c, "Invalid attribute type. Must be one of: "+strings.Join(ValidAttributeTypes, ", "), nil)
		return
	}

	// Create attribute
	attr := models.MetaAttr{
		AttrName: req.AttrName,
		AttrType: req.AttrType,
		AttrDesc: req.AttrDesc,
	}

	if err := database.DB.Create(&attr).Error; err != nil {
		utils.InternalError(c, "Failed to create attribute", err)
		return
	}

	utils.Success(c, attr)
}

func UpdateAttribute(c *gin.Context) {
	attrID := c.Param("id")
	var req AttributeUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request parameters", err)
		return
	}

	var attr models.MetaAttr
	if err := database.DB.First(&attr, attrID).Error; err != nil {
		utils.NotFound(c, "Attribute not found", err)
		return
	}

	attr.AttrName = req.AttrName
	if req.AttrType != nil {
		attr.AttrType = *req.AttrType
	}
	if req.AttrDesc != nil {
		attr.AttrDesc = *req.AttrDesc
	}

	if err := database.DB.Save(&attr).Error; err != nil {
		utils.InternalError(c, "Failed to update attribute", err)
		return
	}

	utils.Success(c, attr)
}

func DeleteAttribute(c *gin.Context) {
	attrID := c.Param("id")

	// 使用事务确保属性和关联关系的删除操作是原子的
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		// 软删除属性
		if err := tx.Model(&models.MetaAttr{}).Where("id = ?", attrID).Update("is_deleted", true).Error; err != nil {
			return err
		}

		// 软删除所有相关的关联关系
		if err := tx.Model(&models.MetaRelation{}).Where("attr_id = ?", attrID).Update("is_deleted", true).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		utils.InternalError(c, "Failed to delete attribute and its relations", err)
		return
	}

	utils.SuccessWithMessage(c, "Attribute and its relations deleted successfully", nil)
}

func GetAttributeList(c *gin.Context) {
	// 获取分页参数
	pageNumStr := c.Query("page_num")
	pageSizeStr := c.Query("page_size")

	// 获取搜索关键词
	keyword := c.Query("keyword")

	var attrs []models.MetaAttr
	var total int64

	// 创建查询构建器
	query := database.DB.Model(&models.MetaAttr{}).Where("is_deleted = ?", false)

	// 如果有搜索关键词，添加搜索条件
	if keyword != "" {
		query = query.Where("attr_name LIKE ? OR attr_desc LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总记录数
	query.Count(&total)

	// 如果没有传分页参数，返回所有数据
	if pageNumStr == "" || pageSizeStr == "" {
		err := query.Order("id DESC").Find(&attrs).Error
		if err != nil {
			utils.InternalError(c, "Failed to get attribute list", err)
			return
		}

		utils.Success(c, gin.H{
			"attributes": attrs,
			"pagination": gin.H{
				"page_num":   1,
				"page_size":  total,
				"total":      total,
				"total_page": 1,
			},
		})
		return
	}

	// 有分页参数时进行分页查询
	pageNum, _ := strconv.Atoi(pageNumStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)

	// 添加排序并使用分页获取数据
	err := utils.Paginate(pageNum, pageSize, query.Order("id DESC")).Find(&attrs).Error
	if err != nil {
		utils.InternalError(c, "Failed to get attribute list", err)
		return
	}

	utils.Success(c, gin.H{
		"attributes": attrs,
		"pagination": gin.H{
			"page_num":   pageNum,
			"page_size":  pageSize,
			"total":      total,
			"total_page": utils.TotalPage(total, pageSize),
		},
	})
}

func GetAttributeInfo(c *gin.Context) {
	attrID := c.Param("id")
	var attr models.MetaAttr
	if err := database.DB.Preload("Events").First(&attr, attrID).Error; err != nil {
		utils.NotFound(c, "Attribute not found", err)
		return
	}

	utils.Success(c, attr)
}

func GetEventAttributes(c *gin.Context) {
	eventID := c.Param("id")

	// 首先检查事件是否存在且未删除
	var event models.MetaEvent
	if err := database.DB.Where("id = ? AND is_deleted = ?", eventID, false).First(&event).Error; err != nil {
		utils.NotFound(c, "Event not found or has been deleted", err)
		return
	}

	// 通过关联表查询属性，只获取未删除的关联和属性
	var attrs []models.MetaAttr
	if err := database.DB.Joins("JOIN meta_relation ON meta_attr.id = meta_relation.attr_id").
		Where("meta_relation.event_id = ? AND meta_relation.is_deleted = ? AND meta_attr.is_deleted = ?",
			eventID, false, false).
		Order("meta_attr.id DESC").
		Find(&attrs).Error; err != nil {
		utils.InternalError(c, "Failed to get event attributes", err)
		return
	}

	utils.Success(c, attrs)
}

// Relation handlers
func CreateRelation(c *gin.Context) {
	var req RelationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, "Invalid request", err)
		return
	}

	// Check if event exists and is not deleted
	var event models.MetaEvent
	if err := database.DB.Where("id = ? AND is_deleted = ?", req.EventID, false).First(&event).Error; err != nil {
		utils.BadRequest(c, "Event not found or has been deleted", err)
		return
	}

	// Check if attribute exists and is not deleted
	var attr models.MetaAttr
	if err := database.DB.Where("id = ? AND is_deleted = ?", req.AttrID, false).First(&attr).Error; err != nil {
		utils.BadRequest(c, "Attribute not found or has been deleted", err)
		return
	}

	// Check if relation already exists
	var existingRelation models.MetaRelation
	err := database.DB.Where("event_id = ? AND attr_id = ?", req.EventID, req.AttrID).First(&existingRelation).Error
	if err == nil {
		utils.BadRequest(c, "Relation already exists", nil)
		return
	} else if err != gorm.ErrRecordNotFound {
		utils.InternalError(c, "Failed to check existing relation", err)
		return
	}

	// Create relation
	relation := models.MetaRelation{
		EventID: req.EventID,
		AttrID:  req.AttrID,
	}

	if err := database.DB.Create(&relation).Error; err != nil {
		utils.InternalError(c, "Failed to create relation", err)
		return
	}

	utils.SuccessWithMessage(c, "Relation created successfully", gin.H{
		"id":       relation.ID,
		"event_id": relation.EventID,
		"attr_id":  relation.AttrID,
	})
}

func GetRelationList(c *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var relations []RelationDetail
	var total int64

	// 创建查询构建器，使用JOIN获取事件名和属性名
	query := database.DB.Table("meta_relation mr").
		Select("mr.id, mr.event_id, mr.attr_id, mr.created_at, me.event_name, ma.attr_name").
		Joins("LEFT JOIN meta_event me ON mr.event_id = me.id").
		Joins("LEFT JOIN meta_attr ma ON mr.attr_id = ma.id").
		Where("mr.is_deleted = ? AND me.is_deleted = ? AND ma.is_deleted = ?", false, false, false).
		Order("mr.id DESC")

	// 获取总记录数
	query.Count(&total)

	// 使用分页方法获取数据
	err := utils.Paginate(pageNum, pageSize, query).Find(&relations).Error
	if err != nil {
		utils.InternalError(c, "Failed to get relation list", err)
		return
	}

	// 格式化创建时间
	for i := range relations {
		relations[i].CreatedStr = relations[i].CreatedAt.Format("2006-01-02 15:04:05")
	}

	utils.Success(c, gin.H{
		"relations": relations,
		"pagination": gin.H{
			"page_num":   pageNum,
			"page_size":  pageSize,
			"total":      total,
			"total_page": utils.TotalPage(total, pageSize),
		},
	})
}

func DeleteRelation(c *gin.Context) {
	relationID := c.Param("id")

	// 检查关联关系是否存在
	var relation models.MetaRelation
	if err := database.DB.First(&relation, relationID).Error; err != nil {
		utils.NotFound(c, "Relation not found", err)
		return
	}

	// 软删除关联关系
	if err := database.DB.Model(&relation).Update("is_deleted", true).Error; err != nil {
		utils.InternalError(c, "Failed to delete relation", err)
		return
	}

	utils.SuccessWithMessage(c, "Relation deleted successfully", nil)
}

// GetAllEventsWithAttributes 获取所有事件及其属性，支持事件搜索和分页
func GetAllEventsWithAttributes(c *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("page_num", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	// 获取搜索关键词
	keyword := c.Query("keyword")

	var events []models.MetaEvent
	var total int64

	// 创建查询构建器
	query := database.DB.Model(&models.MetaEvent{}).
		Preload("Attributes", func(db *gorm.DB) *gorm.DB {
			return db.Where("is_deleted = ?", false).Order("id DESC")
		}).
		Where("meta_event.is_deleted = ?", false)

	// 添加搜索条件
	if keyword != "" {
		query = query.Where("meta_event.event_name LIKE ? OR meta_event.event_desc LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总记录数
	query.Count(&total)

	// 添加排序并使用分页获取数据
	err := utils.Paginate(pageNum, pageSize, query.Order("meta_event.id DESC")).Find(&events).Error
	if err != nil {
		utils.InternalError(c, "Failed to get events with attributes", err)
		return
	}

	utils.Success(c, gin.H{
		"events": events,
		"pagination": gin.H{
			"page_num":   pageNum,
			"page_size":  pageSize,
			"total":      total,
			"total_page": utils.TotalPage(total, pageSize),
		},
	})
}

// GetStatistics 获取统计信息
func GetStatistics(c *gin.Context) {
	var userCount int64
	var eventCount int64
	var attrCount int64
	var relationCount int64

	// 统计用户总数（未删除的）
	if err := database.DB.Model(&models.AdminUser{}).Where("is_deleted = ?", false).Count(&userCount).Error; err != nil {
		utils.InternalError(c, "Failed to count users", err)
		return
	}

	// 统计事件总数（未删除的）
	if err := database.DB.Model(&models.MetaEvent{}).Where("is_deleted = ?", false).Count(&eventCount).Error; err != nil {
		utils.InternalError(c, "Failed to count events", err)
		return
	}

	// 统计属性总数（未删除的）
	if err := database.DB.Model(&models.MetaAttr{}).Where("is_deleted = ?", false).Count(&attrCount).Error; err != nil {
		utils.InternalError(c, "Failed to count attributes", err)
		return
	}

	// 统计关联关系总数（未删除的）
	if err := database.DB.Model(&models.MetaRelation{}).Where("is_deleted = ?", false).Count(&relationCount).Error; err != nil {
		utils.InternalError(c, "Failed to count relations", err)
		return
	}

	utils.Success(c, gin.H{
		"user_count":     userCount,
		"event_count":    eventCount,
		"attr_count":     attrCount,
		"relation_count": relationCount,
	})
}

// GetAttributeTypes 获取属性数据类型列表
func GetAttributeTypes(c *gin.Context) {
	types := []AttributeType{
		{Value: "string", Label: "字符串"},
		{Value: "number", Label: "数值"},
		{Value: "bool", Label: "布尔值"},
		{Value: "time", Label: "日期"},
	}

	utils.Success(c, gin.H{
		"types": types,
	})
}

// isValidAttributeType 检查属性类型是否有效
func isValidAttributeType(attrType string) bool {
	return slices.Contains(ValidAttributeTypes, attrType)
}
