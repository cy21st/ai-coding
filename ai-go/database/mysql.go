package database

import (
	"fmt"
	"log"
	"time"

	"ai-go/config"
	"ai-go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GlobalConfig.MySQL.User,
		config.GlobalConfig.MySQL.Password,
		config.GlobalConfig.MySQL.Host,
		config.GlobalConfig.MySQL.Port,
		config.GlobalConfig.MySQL.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get the underlying *sql.DB instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying *sql.DB: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(config.GlobalConfig.MySQL.MaxIdleConns) // 设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(config.GlobalConfig.MySQL.MaxOpenConns) // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(                                     // 设置连接可复用的最大时间
		time.Duration(config.GlobalConfig.MySQL.ConnMaxLifetime) * time.Minute,
	)

	DB = db.Debug()

	/*
		当你运行程序，GORM 会自动：
		创建表（如果表不存在）
		添加缺失的字段（如果模型中新增了字段）
		修改字段类型（如果模型中的字段类型改变）
		但不会删除现有的字段或表
	*/

	// Auto migrate the schema
	if config.GlobalConfig.MySQL.AutoMigrate {
		log.Println("Starting database auto migration...")
		err = DB.AutoMigrate(
			&models.AdminUser{},
			&models.MetaEvent{},
			&models.MetaAttr{},
			&models.MetaRelation{},
		)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
		log.Println("Database migration completed successfully")
	}
}
