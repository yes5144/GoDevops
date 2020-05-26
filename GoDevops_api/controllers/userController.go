package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yes5144/GoDevops/GoDevops_api/dto"
	"github.com/yes5144/GoDevops/GoDevops_api/models"
	"github.com/yes5144/GoDevops/GoDevops_api/utils"
	"golang.org/x/crypto/bcrypt"
)

// Register xxx
func Register(c *gin.Context) {
	// get post param
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	log.Println("params come from web:", name, telephone, password)

	// check password
	if len(password) < 6 {
		utils.Fail(c, nil, "length of password is less 6")
		return
	}

	// check telephone
	if len(telephone) != 11 {
		utils.Fail(c, nil, "telephone is invalid")
		return
	}

	DB := models.GetDB()
	// check telephone exist
	var user models.User
	if models.IsTelephoneExist(DB, user, telephone) {
		utils.Fail(c, nil, "telephone has already exist msg")
		return
	}

	// hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		utils.Fail(c, nil, "password hashed fail msg")
		return
	}

	// name if not exist
	if len(name) == 0 {
		// name = "slkdjfkjkl"
		name = utils.RandomString(10)
	}
	log.Println("params will be inserted into db:", name, telephone, password)
	user = models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword),
	}
	DB.Create(&user)

	// expected result
	utils.Success(c, nil, "register success msg")
}

// Login xxx
func Login(c *gin.Context) {
	// get param
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	log.Println("params come from web:", name, telephone, password)

	// check telephone
	if len(telephone) != 11 {
		utils.Fail(c, nil, "telephone is invalid")
		return
	}

	// check password
	if len(password) < 6 {
		utils.Fail(c, nil, "length of password can not be less 6")
		return
	}

	// telephone is in db?
	var user models.User
	DB := models.GetDB()
	DB.Where("telephone=?", telephone).First(&user)
	log.Printf("user: %#v", user)
	log.Printf("user-password: %#v", user.Password)

	// compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	log.Println("compare result: ", err)
	if err != nil {
		utils.Fail(c, nil, "password is wrong")
		return
	}

	// generateToken
	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.Fail(c, nil, "generateToken error")
		return
	}
	log.Println(token)

	// expected result
	utils.Success(c, gin.H{"token": token}, "login success msg")

}

// Info xxx
func Info(c *gin.Context) {
	// expected result
	user, _ := c.Get("user")
	utils.Success(c, gin.H{"user": dto.ToUserDto(user.(models.User))}, "info msg")
}
