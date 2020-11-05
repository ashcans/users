package model

// EntityRole 角色
type EntityRole struct {
	RoleId       int64  `gorm:"primaryKey" json:"role_id"`
	RoleName     string `gorm:"type:varchar(100)" json:"role_name"`
	Remark       string `gorm:"type:varchar(100)" json:"remark"`
	CreateUserId int64  `json:"create_user_id"`
	Model
}

func (EntityRole) TableName() string {
	return "role"
}

// EntityUserRole 用户角色对应关系
type EntityUserRole struct {
	ID     int64 `gorm:"primaryKey" json:"id"`
	UserId int64 `gorm:"index:" json:"user_id"`
	RoleId int64 `gorm:"index:" json:"role_id"`
}

func (EntityUserRole) TableName() string {
	return "user_role"
}
