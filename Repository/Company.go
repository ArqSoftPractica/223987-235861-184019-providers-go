package Repository

import (
	"223987-235861-184019-providers/Config"
	"223987-235861-184019-providers/Models"

	_ "github.com/go-sql-driver/mysql"
)

func UpsertCompany(company *Models.Company) (err error) {
	if err = Config.DB.FirstOrCreate(company, Models.Company{ID: company.ID}).Error; err != nil {
		return err
	}
	return nil
}
