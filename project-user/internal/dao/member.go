package dao

import (
	"context"
	"test.com/project-user/internal/database/gorms"
)

type MemberDao struct {
	conn *gorms.GormConn
}

func (m MemberDao) GetMemberByEmail(ctx context.Context, email string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberDao) GetMemberByAccount(ctx context.Context, account string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (m MemberDao) GetMemberByMobile(ctx context.Context, mobile string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewMemberDao() *MemberDao {
	return &MemberDao{
		conn: gorms.New(),
	}
}
