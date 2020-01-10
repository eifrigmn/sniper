package user

import (
	"context"
	"github.com/google/uuid"
	pb "sniper/rpc/user/v1"
	"sniper/util/db"
	"time"
)

// GetUserByName 根据账户名称获取用户账号信息
func (d *Dao) GetUserByName(ctx context.Context, uname string) (*Account, error) {
	query := db.SQLSelect("tb_account", `select id, account, email, pswd from tb_account where account = ? and deleted_at is NULL`)
	userInfoRow := d.db.QueryRowContext(ctx, query, uname)
	var userInfo Account
	err := userInfoRow.Scan(&userInfo.ID, &userInfo.Account, &userInfo.Mail, &userInfo.Password)
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

// CreateUser 用户信息写入DB
func (d *Dao) CreateUser(ctx context.Context, data *pb.RegistReq) error {
	uid := uuid.Must(uuid.NewRandom()).String()
	query := db.SQLInsert("tb_account", "insert into tb_account (id, account, email, pswd, created_at) values (?,?,?,?,?)")
	_, err := d.db.ExecContext(ctx, query, uid, data.GetAccount(), data.GetMail(), data.GetPassword(), time.Now())
	return err
}
