package initialize

import (
	"fmt"
	"jsonrpcmicro/internal/auth/config"
	"jsonrpcmicro/internal/auth/model"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := config.Conf.DataSource
	if m.Database == "" {
		return nil
	}
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Address + ":" + m.Port + ")/" + m.Database + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		return nil
	} else {
		return db
	}
}

func InitMysql(db *gorm.DB) {
	err := db.AutoMigrate(
		model.AuthUser{},
		model.AuthGroup{},
		model.AuthApi{},
		model.Authority{},
	)
	if err != nil {
		fmt.Println("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
