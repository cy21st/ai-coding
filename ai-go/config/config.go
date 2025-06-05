package config

type Config struct {
	MySQL MySQLConfig
	Redis RedisConfig
	JWT   JWTConfig
}

type MySQLConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` // in minutes
	AutoMigrate     bool   `mapstructure:"auto_migrate"`      // 是否自动迁移数据库
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type JWTConfig struct {
	SecretKey string
	Expire    int // in hours
}

var GlobalConfig = Config{
	MySQL: MySQLConfig{
		Host:            "localhost",
		Port:            3306,
		User:            "root",
		Password:        "root",
		Database:        "meta_data",
		MaxIdleConns:    10,    // 设置空闲连接池中的最大连接数
		MaxOpenConns:    100,   // 设置打开数据库连接的最大数量
		ConnMaxLifetime: 60,    // 设置连接可复用的最大时间（分钟）
		AutoMigrate:     false, // 默认不开启自动迁移
	},
	Redis: RedisConfig{
		Host:     "localhost",
		Port:     6379,
		Password: "",
		DB:       0,
	},
	JWT: JWTConfig{
		SecretKey: "ai-code",
		Expire:    24, // 24 hours
	},
}
