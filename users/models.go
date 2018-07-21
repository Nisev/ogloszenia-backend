package users

import (
	"github.com/jinzhu/gorm"
	"github.com/nifrez/ogloszenia/common"
	"github.com/rs/zerolog/log"
)

type User struct {
	gorm.Model
	AccountId string `gorm:"type:varchar(255);" json:"accountId"`
	Username  string `gorm:"type:varchar(255);" json:"username"`
}

func AutoMigrate(){
	db := common.GetDB()

	db.AutoMigrate(&User{})
	log.Info().Msg("Auto migrating finished.")
}
