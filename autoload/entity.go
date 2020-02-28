package autoload

import (
	"log"

	"graphi/datastore"
	"graphi/domain/model"
)

func Load() {
	db, err := datastore.NewMyqlDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&model.Invoice{}, &model.Balance{}, &model.Event{}, &model.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().DropTableIfExists(&model.UserRole{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().AutoMigrate(&model.User{}, &model.Event{}, &model.UserRole{}, &model.Balance{}, &model.Invoice{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Event{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.User{}).AddForeignKey("user_role_id", "user_roles(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Balance{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&model.Invoice{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().Model(&model.Invoice{}).AddForeignKey("event_id", "events(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range user_roles{
		err = db.Debug().Model(&model.UserRole{}).Create(&user_roles[i]).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	for i, _ := range users{
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatal(err)
		}
	}
}
