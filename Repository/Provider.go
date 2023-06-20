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

func GetAllProviders(providers *[]Models.Provider, companyId string) (err error) {
	if err = Config.DB.Find(providers).Where("id = ?", companyId).Error; err != nil {
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

func GetProviderByID(provider *Models.Provider, id string, companyId string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(provider).Where("id = ? AND company_id = ?", id, companyId).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProvider(provider *Models.Provider, id string) (err error) {
	var existingProvider Models.Provider
	if err = Config.DB.Where("id = ? AND company_id = ?", id, provider.CompanyId).First(&existingProvider).Error; err != nil {
		return err
	}

	fmt.Println(provider)
	Config.DB.Save(provider)
	return nil
}

func DeactivateProvider(provider *Models.Provider, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(provider).Where("id = ? AND company_id = ?", id, provider.CompanyId).Error; err != nil {
		return err
	}
	provider.IsActive = false
	fmt.Println(provider)
	Config.DB.Save(provider)
	return nil
}
