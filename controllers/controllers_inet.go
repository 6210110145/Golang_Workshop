package controllers

import (
	"errors"
	// "fmt"
	"strings"

	"go-fiber-test/database"
	m "go-fiber-test/models"
	"log"
	"regexp"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func HelloV2(c *fiber.Ctx) error {
	return c.SendString("Hello V2")
}

func BodyParserTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	}

	log.Println(p.Name)
	log.Println(p.Pass)

	str := p.Name + p.Pass
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {
	str := "hello ==>" + c.Params("name")

	return c.JSON(str)
}

func QueryTest(c *fiber.Ctx) error {
	a := c.Query("search")
	str := "my search is " + a
	return c.JSON(str)
}

func ValidTest(c *fiber.Ctx) error {
	//Connect to database

	user := new(m.User)
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	validate := validator.New()
	errors := validate.Struct(user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.Error())
	}
	return c.JSON(user)
}

func Factorial(c *fiber.Ctx) error {
	x := c.Params("num")
	x1, _ := strconv.Atoi(x)

	var result int
	result = 1

	if x1 < 0 {
		return errors.New("the number must be more then 0")
	}

	if x1 == 0 {
		str := x + "! = " + strconv.Itoa(result)
		return c.JSON(str)
	} else {
		for i := x1; i > 0; i-- {
			result *= i
		}
	}
	str := x + "! = " + strconv.Itoa(result)
	return c.JSON(str)
}

func Acii(c *fiber.Ctx) error {
	name := c.Query("tax_id")
	acii := []string{}

	for _, v := range name {
		acii = append(acii, strconv.Itoa(int(v)))
	}
	return c.JSON(strings.Join(acii, " "))
}

func ValidateRegister(c *fiber.Ctx) error {
	newCompany := new(m.Company01)
	validate := validator.New()

	validate.RegisterValidation("customName", ValidateCustomName)
	validate.RegisterValidation("customWeb", ValidateCustomWeb)

	if err := c.BodyParser(&newCompany); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errorValidate := validate.Struct(newCompany)
	// fmt.Println(errorValidate)
	if errorValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidate.Error(),
		})
	}

	return c.JSON(newCompany)
}

func ValidateCustomName(fl validator.FieldLevel) bool {
	requiredName := `^[a-z,A-Z,0-9,_,-]+$`
	matched, _ := regexp.MatchString(requiredName, fl.Field().String())

	return matched
}

func ValidateCustomWeb(fl validator.FieldLevel) bool {
	requiredWeb := `^[http://,https://,a-z,0-9,-]+$`
	matched, _ := regexp.MatchString(requiredWeb, fl.Field().String())

	return matched
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog []m.Dogs
	search := strings.TrimSpace(c.Query("search"))

	result := db.Find(&dog, "dog_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&dog)
}

// exercise 7.0.2
func GetDeleteDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)
	
	return c.Status(200).JSON(dogs)
}

// exercise 7.1
func GetSomeDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Where("dog_id BETWEEN 51 AND 99").Find(&dogs)

	return c.Status(200).JSON(dogs)
}

func AddDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)

	return c.Status(201).JSON(dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)

	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

// exercise 7.2
func GetDogJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs
	var red, green, pink, nocolor int

	db.Find(&dogs)

	var dataResult []m.DogsRes
	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			red++
		}else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			green++
		}else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			pink++
		}else {
			typeStr = "no color"
			nocolor++
		}

		d := m.DogsRes{
			Name: v.Name,
			DogID: v.DogID,
			Type: typeStr,
		}

		dataResult = append(dataResult, d)
	}

	r := m.ResultData{
		Count: len(dogs),
		Data: dataResult,
		Name: "goleng-test",
		Red: red,
		Green: green,
		Pink: pink,
		Nocolor: nocolor,
	}

	return c.Status(200).JSON(r)
}

//exercise 7.0.1
func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var newCompany m.Company

	if err := c.BodyParser(&newCompany); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	validate := validator.New()

	errorValidate := validate.Struct(newCompany)
	if errorValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidate.Error(),
		})
	}

	result := db.Where("name = ?", string(newCompany.Name)).Find(&newCompany)

	if result.RowsAffected != 0 {
		return c.Status(401).SendString("name is required")
	}

	db.Create(&newCompany)

	return c.Status(201).JSON(newCompany)
}

func GetAllCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company

	db.Find(&company)

	return c.Status(200).JSON(company)
}

func GetNameCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company []m.Company
	search := strings.TrimSpace(c.Query("search"))

	result := db.Find(&company, "name = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&company)

	return c.Status(200).JSON(fiber.Map{
		"data": company,
		"message": "update success",
	})
}

func DeleteCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).SendString("delete success")
}

//project
func CreateProfile(c *fiber.Ctx) error {
	db := database.DBConn
	var newprofile m.Profile

	if err := c.BodyParser(&newprofile); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&newprofile)

	return c.Status(201).JSON(newprofile)
}

func ShowProfiles(c *fiber.Ctx) error {
	db := database.DBConn
	var profile []m.Profile

	db.Find(&profile)

	return c.Status(200).JSON(profile)
}