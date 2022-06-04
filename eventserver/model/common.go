package model

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"memo-RPC/eventserver/conf"
)

var DB *gorm.DB

func Setup() {
	var err error

	// pass conf to dsn, meet the problem that there is not
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conf.Cfg.Database.Host,
		conf.Cfg.Database.User,
		conf.Cfg.Database.Password,
		conf.Cfg.Database.Name,
		conf.Cfg.Database.Port,
		conf.Cfg.Database.SSLMode,
		conf.Cfg.Database.TimeZone,
	)

	// open the database and buffer the conf
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 迁移时禁用外键
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Cfg.Database.TablePrefix, // set the prefix name of table
			SingularTable: true,                          // use singular table by default
		},
		Logger: logger.Default.LogMode(logger.Info), // set log mode
	})

	// some init set of database
	pgSQL, err := DB.DB()
	if err != nil {
		log.Panicln("db.DB() err: ", err)
	}
	pgSQL.SetMaxIdleConns(conf.Cfg.Database.SetMaxIdleConns)       // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	pgSQL.SetMaxOpenConns(conf.Cfg.Database.SetMaxOpenConns)       // SetMaxOpenConns 设置打开数据库连接的最大数量
	pgSQL.SetConnMaxLifetime(conf.Cfg.Database.SetConnMaxLifetime) // SetConnMaxLifetime 设置了连接可复用的最大时间

	// set auto migrate
	DB.AutoMigrate(
		&Event{},
	)
}
