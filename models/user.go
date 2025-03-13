package models

import "time"

// UserRole 枚举定义
type UserRole string

const (
	ProcurementUser     UserRole = "procurement"
	EquipmentManagerUser UserRole = "equipment_manager"
	SuperUser           UserRole = "superuser"
	OperatorUser        UserRole = "operator"
)

// BaseUser 公共用户结构
type BaseUser struct {
	ID       uint     `gorm:"primaryKey"`
	Username string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Profile  string   `gorm:"unique"`
	Role     UserRole `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 采购用户
type Procurement struct {
	BaseUser
}

// 设备管理用户
type EquipmentManager struct {
	BaseUser
}

// 超级用户
type Super struct {
	BaseUser
}

// 操作员用户
type Operator struct {
	BaseUser
}