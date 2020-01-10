package user

import (
	"context"
	"sniper/util/db"
)

const (
	// DBName 数据库名称
	DBName = "default"
	// TbAccount 用户信息表名称
	TbAccount = "tb_account"
)

// Dao 数据层
type Dao struct {
	db *db.DB
}

// New 新建数据库连接
func New(ctx context.Context, dbName string) *Dao {
	conn := db.Get(ctx, DBName)
	return &Dao{db: conn}
}

// Account 用户账号
type Account struct {
	ID       string
	Account  string
	Password string
	Mail     string
}
