package Repository

import (
	"223987-235861-184019-providers/Config"
	"223987-235861-184019-providers/Models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateCompany(company *Models.Provider) (err error) {
	if err = Config.DB.Create(company).Error; err != nil {
		return err
	}
	return nil
}
