package svc_user

import "time"

const (
	CredKindMPOpenID = "mp-open-id"
)

type User struct {
	ID        int64     `gorm:"column:id;primary_key;auto_increment:false"`
	Nickname  string    `gorm:"column:nickname;type:varchar(100)"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type Auth struct {
	UserID    int64     `gorm:"column:user_id;index:idx_auth_user_id"`
	Kind      string    `gorm:"column:kind;type:varchar(100);not null;unique_index:idx_auth_source"`
	Name      string    `gorm:"column:name;type:varchar(100);not null;unique_index:idx_auth_source"`
	Secret    string    `gorm:"column:secret;type:text"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
