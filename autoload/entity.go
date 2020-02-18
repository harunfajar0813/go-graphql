package autoload

import (
	"log"

	"graphi/domain/model"
	"graphi/datastore"
)

func Load() {
	db, err := datastore.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&model.Event{}, &model.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&model.User{}, &model.Event{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Event{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
}
