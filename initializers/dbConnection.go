package initializers
import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDb (){

		dsn := "host=localhost user=postgres password=yourpassword dbname=trekkers port=5432"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("Failed to connect to the database!")
		}
	
		db.AutoMigrate()
		DB = db
	
	
}