package handlers

import (
	"ai-go/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Test(c *gin.Context) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		utils.InternalError(c, "Failed to hash password", err)
		return
	}
	utils.Success(c, string(hashedPassword))
}

func Test2(c *gin.Context) {

	// var relations []models.MetaRelation

	// database.DB.Preload("Event").Preload("Attr").
	// 	Joins("LEFT JOIN meta_event ON meta_relation.event_id = meta_event.id").
	// 	Joins("LEFT JOIN meta_attr ON meta_relation.attr_id = meta_attr.id").
	// 	Where("meta_relation.is_deleted = ?", false).
	// 	Where("meta_event.is_deleted = ?", false).
	// 	Where("meta_attr.is_deleted = ?", false).
	// 	Order("meta_relation.id DESC").
	// 	Limit(10).
	// 	Find(&relations)

	// database.DB.Preload("Event", "is_deleted = ?", false).
	// 	Preload("Attr", "is_deleted = ?", false).
	// 	Where("meta_relation.is_deleted = ?", false).
	// 	Order("meta_relation.id DESC").
	// 	Limit(10).
	// 	Find(&relations)

	// database.DB.Preload("Event", func(db *gorm.DB) *gorm.DB {
	// 	return db.Select("id, event_name").Where("is_deleted = ?", false)
	// }).
	// 	Preload("Attr", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("id, attr_name").Where("is_deleted = ?", false)
	// 	}).
	// 	Where("is_deleted = ?", false).
	// 	Order("id DESC").
	// 	Limit(10).
	// 	Find(&relations)

	utils.Success(c, "test2")
}
