package model

// EntityRole 角色
type EntityRole struct {
	RoleId uint64 `gorm:"primaryKey" json:"role_id"`
	RoleName string `gorm:"type:varchar(255)" json:"role_name"`
	Remark string `gorm:"type:varchar(255)" json:"remark"`
	CreateUserId uint64 `json:"create_user_id"`
	Model
}

// EntityUserRole 用户角色对应关系
type EntityUserRole struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
	UserId uint64 `gorm:"index:" json:"user_id"`
	RoleId uint64 `gorm:"index:" json:"role_id"`
}