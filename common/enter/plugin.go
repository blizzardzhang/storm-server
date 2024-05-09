package enter

import (
	"fmt"
	"gorm.io/gorm"
	"storm-server/common/uniqueid"
)

func Initialize(db *gorm.DB) (err error) {
	// 创建字段的时候雪花算法生成id
	e := db.Callback().Create().Before("gorm:create").Replace("id", func(db *gorm.DB) {
		id := uniqueid.GenIdStr()
		db.Statement.SetColumn("id", fmt.Sprintf("%d", id))
		userId := db.Statement.Context.Value("userId")
		db.Statement.SetColumn("created_by", userId)
	})
	if e != nil {
		return e
	}
	return

}
