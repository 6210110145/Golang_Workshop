package controllers

import (
	"errors"
	"fmt"
	"strings"

	// "fmt"
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

	// if err != nil {
	// 	// return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
	// 	return errors.New("empty parameter")
	// }

	if x1 < 0 {
		return errors.New("the number must be more then 0")
	}

	if x1 == 0 {
		str := x + "! = " + strconv.Itoa(result)
		return c.JSON(str)
	}else {
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
	newCompany := new(m.Company)
	validate := validator.New()

	validate.RegisterValidation("customName", ValidateCustomName)
	validate.RegisterValidation("customWeb", ValidateCustomWeb)

	if err := c.BodyParser(&newCompany); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errorValidate := validate.Struct(newCompany)
	fmt.Println(errorValidate)
	if errorValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": errorValidate.Error(),
		})
	}

	return c.JSON(newCompany)
}

func ValidateCustomName(fl validator.FieldLevel) bool  {
	requiredName := `^[a-z,A-Z,0-9,_,-]+$`
	matched, _ := regexp.MatchString(requiredName, fl.Field().String())

	return matched
}

func ValidateCustomWeb(fl validator.FieldLevel) bool {
	requiredWeb := `^[http://,https://,a-z,0-9,-]+$`
	matched, _ := regexp.MatchString(requiredWeb, fl.Field().String())

	return matched
}