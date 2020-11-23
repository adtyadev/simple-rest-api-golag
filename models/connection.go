package models

import(
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)
var DB *gorm.DB
func Connection(){
 //db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
 dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
 db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
 if err != nil {
   panic("failed to connect database")
 }
  // Migrate the schema
  db.AutoMigrate(&Book{})
 DB = db
}