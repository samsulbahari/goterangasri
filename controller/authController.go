package controller

import (
	"goterangasri/config"
	_ "goterangasri/config"
	"goterangasri/helper"
	"goterangasri/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Nama             string `json:"nama" binding:"required"`
	Username         string `json:"username" binding:"required"`
	Password         string `json:"password" binding:"required"`
	Confirm_password string `json:"confirm_password" binding:"eqfield=Password,required"`
}

func Registerinsert(c *gin.Context) {
	var user UserInput
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.Validatorerror(c, err)
	} else {
		db, _ := config.SetDatabase()
		getuser := model.User{}
		err := db.Where("username = ?", user.Username).First(&getuser).Error
		if err != nil {
			hash, _ := helper.HashPassword(user.Password)
			getuser.Nama = user.Nama
			getuser.Username = user.Username
			getuser.Password = hash

			db.Create(&getuser)

			c.JSON(http.StatusCreated, gin.H{
				"code":    http.StatusOK,
				"message": "successfully",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    http.StatusOK,
				"message": "failed username tersedia",
			})

		}

	}

}
