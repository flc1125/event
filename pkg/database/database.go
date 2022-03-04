package database

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Driver   string
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
	Options  *gorm.Config
}

type Connection struct {
	Name   string
	DB     *gorm.DB
	Config *Config
}

var Connections = map[string]*Connection{}
var Configs = map[string]*Config{}

// Configure 配置数据库连接
func Configure(name string, config *Config) {
	Configs[name] = config
}

// Connect 创建个 DB 连接
func Connect(name string) *Connection {
	// 如果存在连接，则直接获取
	if db, flag := Connections[name]; flag {
		return db
	}

	// 如果不存在，则创建连接
	return Create(name)
}

// Create 创建个连接
func Create(name string) *Connection {
	// 获取配置
	config, flag := Configs[name]
	if !flag {
		panic(errors.New("不存在的[" + name + "]连接"))
	}

	// 默认值
	connection := &Connection{
		Name:   name,
		Config: config,
	}
	var err error

	// 识别驱动进行 DB 连接
	switch config.Driver {
	case "mysql":
		connection.DB, err = createMySqlDriver(config)
	default:
		panic(errors.New("不支持该[" + config.Driver + "]驱动"))
	}

	if err != nil {
		panic(err)
	}

	Connections[name] = connection

	return Connections[name]
}

// 创建 DB 驱动
func createMySqlDriver(config *Config) (*gorm.DB, error) {
	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=" + config.Charset + "&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
