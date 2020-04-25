package database

import (
	"github.com/iamhabbeboy/api/model"
)

//GameSeeder for migration
func GameSeeder() {

	db := Connect()

	db.DropTable(&model.Game{})

	db.AutoMigrate(&model.Game{})

	games := []model.Game{
		{
			Question:  "Who is this",
			Answer:  "prosper",
			CategoryID: 1,
			Pictures: []model.Picture{
				{
					URL: "https://pbs.twimg.com/media/DqlqiJbWwAICl8C.jpg",
					Size: "100,300,500,1000",
				},
				{
					URL: "https://pbs.twimg.com/media/DqlqiJbWwAICl8C.jpg",
					Size: "100,300,500,1000",
				},
			},
		},
		{
			Question:  "Where is this place around shagamu",
			Answer:  "sabo",
			CategoryID: 2,
			Pictures: []model.Picture{
				{
					URL: "https://pbs.twimg.com/media/DqlqiJbWwAICl8C.jpg",
					Size: "100,300,500,1000",
				},
				{
					URL: "https://pbs.twimg.com/media/DqlqiJbWwAICl8C.jpg",
					Size: "100,300,500,1000",
				},
			},
		},
	}

	for _, game := range games {
		db.Create(&game)
	}
}

//PictureSeeder migration
func PictureSeeder() {
	db := Connect()

	db.DropTable(&model.Picture{})

	db.AutoMigrate(&model.Picture{})

	pictures := []model.Picture{
		{
			URL: "https://pbs.twimg.com/media/DqlqiJbWwAICl8C.jpg",
			Size: "100,300,500,1000",
			GameID: 1,
		},
		{
				URL: "https://pbs.twimg.com/media/DqlqiJbWwAICl8C.jpg",
				Size: "100,300,500,1000",
				GameID: 2,
		},
	}

	for _, picture := range pictures {
		db.Create(&picture)
	}
}

//CategorySeeder migration
func CategorySeeder() {
	db := Connect()

	db.DropTable(&model.Category{})

	db.AutoMigrate(&model.Category{})

	categories := []model.Category{
		{
			Title: "People",
			Description: "black, white, hispanic,etc",
		},
		{
			Title: "Places",
			Description: "Country, State, City,etc",
		},
	}
	for _, category := range categories {
		db.Create(&category)
	}
}