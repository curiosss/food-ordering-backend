package postgresql

import (
	"fmt"
	"template2/internal/domain/entities"

	"gorm.io/gorm"
)

func SyncDb(db *gorm.DB) error {
	fmt.Println("syncing database")
	return db.AutoMigrate(&entities.User{}, &entities.Department{})
}

func SeedConstants(db *gorm.DB) {
	// sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Table("users").Where("id", user_id).Delete(&user_id)
	// })
	// fmt.Println(sql)
	fmt.Println("Seeding departments...")
	users := []entities.Department{}
	res := db.Find(&users)
	if res.Error == nil && res.RowsAffected < 4 {
		fmt.Println(res.RowsAffected)
		result := db.Create(&entities.Department{Id: 1, Name: "Restaurants", IsEnabled: true, ImageUrl: "/public/images/restaurant.jpeg"})
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		result = db.Create(&entities.Department{Id: 2, Name: "Hotels", IsEnabled: false, ImageUrl: "/public/images/hotels.jpeg"})
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		result = db.Create(&entities.Department{Id: 3, Name: "Apartments & Offices", IsEnabled: false, ImageUrl: "/public/images/houses.jpeg"})
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		result = db.Create(&entities.Department{Id: 4, Name: "Services", IsEnabled: false, ImageUrl: "/public/images/services.jpeg"})
		if result.Error != nil {
			fmt.Println(result.Error)
		}
	} else {
		fmt.Println(res.Error)
	}

}
