package postgresql

import (
	"fmt"
	"template2/internal/domain/entities"

	// _ "github.com/golang-migrate/migrate/v4/source/file"
	// _ "github.com/lib/pq"
	"gorm.io/gorm"
)

func SyncDb(db *gorm.DB) error {
	fmt.Println("syncing database")
	return db.AutoMigrate(
		&entities.User{},
		&entities.Roles{},
		&entities.UserRoles{},
		&entities.Department{},
	)
	// dbb, err := sql.Open("postgres", "host=localhost user=postgres password=hellothere dbname=test_erp port=5432 sslmode=disable")
	// driver, err := postgres.WithInstance(dbb, &postgres.Config{})
	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file:///migrations",
	// 	"postgres", driver)
	// m.Up()
	// return err
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
