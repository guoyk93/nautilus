package svc_user

import (
	"context"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"go.guoyk.net/nrpc/v2"
	"nautilus/svc/svc_id"
	"strconv"
)

type GetQuery struct {
	MPOpenID string `json:"mp_open_id"`
}

type GetResp struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
}

type ServiceOptions struct {
	DB     *gorm.DB
	Client *nrpc.Client
}

type UserService struct {
	db       *gorm.DB
	idClient *svc_id.Client
}

func NewService(opts ServiceOptions) *UserService {
	return &UserService{
		db:       opts.DB,
		idClient: svc_id.NewClient(opts.Client),
	}
}

func (s *UserService) HealthCheck(ctx context.Context) error {
	return s.db.Exec("SELECT VERSION();").Error
}

func (s *UserService) Get(ctx context.Context, req *GetQuery) (out GetResp, err error) {
	if len(req.MPOpenID) == 0 {
		err = nrpc.Solid(errors.New("missing argument mp_open_id"))
		return
	}
	var newUserID int64
	if newUserID, err = s.idClient.NextOne(ctx); err != nil {
		return
	}
	var auth Auth
	if err = s.db.Where(Auth{Kind: CredKindMPOpenID, Name: req.MPOpenID}).Attrs(Auth{UserID: newUserID}).FirstOrCreate(&auth).Error; err != nil {
		return
	}
	var user User
	if err = s.db.Where(User{ID: auth.UserID}).Attrs(User{Nickname: fmt.Sprintf("用户(%d)", auth.UserID)}).FirstOrCreate(&user).Error; err != nil {
		return
	}
	out.ID = strconv.FormatInt(user.ID, 10)
	out.Nickname = user.Nickname
	return
}
