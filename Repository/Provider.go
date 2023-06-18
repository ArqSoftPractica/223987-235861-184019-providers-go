package Repository

import (
	"223987-235861-184019-providers/Config"
	"223987-235861-184019-providers/Models"
	"223987-235861-184019-providers/Util"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func GetAllProviders(providers *[]Models.Provider) (err error) {
	if err = Config.DB.Find(providers).Error; err != nil {
		return err
	}
	return nil
}

func validateEmail(email string) bool {
	// Regular expression for email syntax validation
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, email)
	return match
}

func validateModel(m *Models.Provider) error {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return err
	}
	return nil
}

func CreateProvider(provider *Models.Provider) (err error) {
	provider.ID = uuid.New()
	err = validateModel(provider)

	if err != nil {
		return err
	}

	if validateEmail(provider.Email) {
		if err = Config.DB.Create(provider).Error; err != nil {
			return err
		}
		return nil
	} else {
		return &Util.CustomError{
			Message: "Email is invalid",
		}
	}
}

func GetProviderByID(provider *Models.Provider, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(provider).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProvider(provider *Models.Provider, id string) (err error) {
	fmt.Println(provider)
	Config.DB.Save(provider)
	return nil
}

func DeleteProvider(provider *Models.Provider, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(provider)
	return nil
}
