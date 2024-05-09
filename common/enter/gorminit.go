package enter

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func InitGorm(MysqlDatasource string, tablePrefix string, models []interface{}) *gorm.DB {
	db, err := gorm.Open(mysql.Open(MysqlDatasource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, //设置表前缀
			SingularTable: true,        //设置单数表名
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect to database success")

	// Migrate the schema
	if err = db.AutoMigrate(models...); err != nil {
		log.Fatalf("auto migrate tables failed: %s", err.Error())
	}
	fmt.Println("Migrate tables success")
	return db
}
