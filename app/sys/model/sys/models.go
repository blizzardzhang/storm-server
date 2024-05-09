package sys

import (
	"gorm.io/gorm"
	"storm-server/common/uniqueid"
	"time"
)

var Models = []interface{}{
	&User{},
	&Client{},
	&Department{},
	&Role{},
	&Permission{},
	&RolePermission{},
	&UserRole{},
}

type Model struct {
	Id        string         `gorm:"size:50;primaryKey;autoIncrement:false" json:"id"`
	CreateBy  string         `gorm:"size:50;" json:"createBy" form:"createBy"`
	UpdateBy  string         `gorm:"size:50;" json:"updateBy" form:"updateBy"`
	CreateAt  time.Time      `json:"CreateAt" form:"createAt"`
	UpdateAt  time.Time      `json:"UpdateAt" form:"updateAt"`
	Status    int            `gorm:"not null;default:0" json:"status" form:"status"` // 状态
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate 定义一个钩子，在创建之前执行自动插入字段
func (model Model) BeforeCreate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("Id", uniqueid.GenIdStr())
	db.Statement.SetColumn("CreateAt", time.Now())
	db.Statement.SetColumn("UpdateAt", time.Now())
	return
}

// BeforeUpdate 定义一个钩子，在更新之前执行自动更新字段
func (model Model) BeforeUpdate(db *gorm.DB) (err error) {
	db.Statement.SetColumn("UpdateAt", time.Now())
	return
}

// User user
type User struct {
	ClientId     string `gorm:"size:32;" json:"ClientId" form:"clientId"`
	DepartmentId string `gorm:"size:50;" json:"departmentId" form:"departmentId"`
	Account      string `gorm:"size:32;unique;" json:"account" form:"account"`
	Name         string `gorm:"size:50;" json:"name" form:"name"`
	Nickname     string `gorm:"size:50;" json:"nickname" form:"nickname"`
	Password     string `gorm:"size:100" json:"password" form:"password"`
	Phone        string `gorm:"size:20;" json:"phone" form:"phone"`
	Model
}

// Client app
type Client struct {
	Name                  string `gorm:"size:50;" json:"name" form:"name"`
	ClientId              string `gorm:"size:50;" json:"clientId" form:"clientId"`
	Key                   string `gorm:"size:100;" json:"key" form:"key"`
	Secret                string `gorm:"size:100;" json:"secret" form:"secret"`
	GrantType             string `gorm:"size:200;" json:"grantType" form:"grantType"`
	RedirectUri           string `gorm:"size:200;" json:"redirectUri" form:"redirectUri"`
	AdditionalInformation string `gorm:"size:200;" json:"additionalInformation" form:"additionalInformation"`
	AccessTokenValidity   int64  `gorm:"not null" json:"accessTokenValidity" form:"accessTokenValidity"`
	RefreshTokenValidity  int64  `gorm:"not null" json:"refreshTokenValidity" form:"refreshTokenValidity"`
	Model
}

// Department dept
type Department struct {
	ParentId  string `gorm:"size:50;index:idx_dept_parent_id" json:"parentId" form:"parentId"`
	Ancestors string `gorm:"size:1000;" json:"ancestors" form:"ancestors"`
	Name      string `gorm:"size:50;" json:"name" form:"name"`
	Sort      int    `gorm:"type:int(11);" json:"sort" form:"sort"`
	Model
}

// Role role
type Role struct {
	Model
	Type   int    `gorm:"not null;default:1" json:"type" form:"type"` // 角色类型（0：系统角色、1：自定义角色）
	Name   string `gorm:"size:50" json:"name" form:"name"`            // 角色名称
	Code   string `gorm:"unique;size:64" json:"code" form:"code"`     // 角色编码
	Sort   int    `json:"sort" form:"sort"`                           // 排序
	Remark string `gorm:"size:256" json:"remark" form:"remark"`       // 备注
}

// Permission perm
type Permission struct {
	ParentId  string `json:"parentId" form:"parentId"`          // 上级菜单
	Name      string `gorm:"size:50" json:"name" form:"name"`   // 名称
	Code      string `gorm:"size:50" json:"title" form:"title"` // 标题
	Component string `gorm:"size:100" json:"component" form:"component"`
	Icon      string `gorm:"size:100" json:"icon" form:"icon"`           // ICON
	Path      string `gorm:"size:100" json:"path" form:"path"`           // 路径
	Sort      int    `gorm:"not null;default:0" json:"sort" form:"sort"` // 排序
	Category  string `gorm:"size:64" json:"category" form:"category"`
	Model
}

// UserRole user_role
type UserRole struct {
	UserId string `gorm:"size:50,uniqueIndex:idx_user_role" json:"userId" form:"userId"`
	RoleId string `gorm:"size:50,uniqueIndex:idx_user_role" json:"roleId" form:"roleId"`
	Model
}

// RolePermission role_permission
type RolePermission struct {
	RoleId       string `gorm:"size:50;uniqueIndex:idx_role_permission" json:"roleId" form:"roleId"`
	PermissionId string `gorm:"size:50;uniqueIndex:idx_role_permission" json:"permissionId" form:"permissionId"`
	Model
}
