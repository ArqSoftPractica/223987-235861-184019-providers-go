package Controllers

import (
	"223987-235861-184019-providers/Models"
	"223987-235861-184019-providers/Repository"
	"223987-235861-184019-providers/Util"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ToCustomError(err error) *Util.CustomError {
	if customErr, ok := err.(*Util.CustomError); ok {
		return customErr
	}
	return &Util.CustomError{
		Message: err.Error(),
	}
}

func GetCompanyIdFromContext(c *gin.Context) (string, error) {
	user := c.Request.Context().Value("user")
	if user == nil {
		return "", &Util.CustomError{
			Message: "No user in Auth",
		}
	}

	claims, ok := user.(jwt.MapClaims)
	if !ok {
		return "", &Util.CustomError{
			Message: "Incorrect user type in auth",
		}
	}

	companyId, ok := claims["companyId"].(string)
	if !ok {
		return "", &Util.CustomError{
			Message: "Incorrect user type in auth",
		}
	}

	log.Printf("CompanyId: %s", companyId)
	return companyId, nil
}

func GetProviders(c *gin.Context) {
	companyId, err := GetCompanyIdFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var providers []Models.Provider
	err = Repository.GetAllProviders(&providers, companyId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, ToCustomError(err))
		return
	}

	c.JSON(http.StatusOK, providers)
}

func CreateProvider(c *gin.Context) {
	companyId, err := GetCompanyIdFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	var provider Models.Provider
	c.BindJSON(&provider)
	uuidObj, errParse := uuid.Parse(companyId)
	if errParse != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
	} else {
		provider.CompanyId = uuidObj
		err := Repository.CreateProvider(&provider)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
		} else {
			c.JSON(http.StatusOK, provider)
		}
	}
}

func GetProviderByID(c *gin.Context) {
	companyId, err := GetCompanyIdFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	} else {
		id := c.Params.ByName("id")
		var provider Models.Provider
		err := Repository.GetProviderByID(&provider, id, companyId)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, ToCustomError(err))
		} else {
			c.JSON(http.StatusOK, provider)
		}
	}
}

func UpdateProvider(c *gin.Context) {
	companyId, err := GetCompanyIdFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	} else {
		id := c.Params.ByName("id")
		providerId, errParse := uuid.Parse(id)
		if errParse != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
		} else {
			companyUuid, errParse := uuid.Parse(companyId)
			if errParse != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
			} else {
				var provider Models.Provider
				err := Repository.GetProviderByID(&provider, id, companyId)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusNotFound, ToCustomError(err))
				} else {
					c.BindJSON(&provider)
					provider.CompanyId = companyUuid
					provider.ID = providerId
					err = Repository.UpdateProvider(&provider, id)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
					} else {
						c.JSON(http.StatusOK, provider)
					}
				}
			}
		}
	}
}

func DeleteProvider(c *gin.Context) {
	companyId, err := GetCompanyIdFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	} else {
		id := c.Params.ByName("id")
		providerId, errParse := uuid.Parse(id)
		if errParse != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
		} else {
			companyUuid, errParse := uuid.Parse(companyId)
			if errParse != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
			} else {
				var provider Models.Provider
				err := Repository.GetProviderByID(&provider, id, companyId)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusNotFound, ToCustomError(err))
				} else {
					c.BindJSON(&provider)
					provider.CompanyId = companyUuid
					provider.ID = providerId
					err = Repository.DeactivateProvider(&provider, id)
					if err != nil {
						c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
					} else {
						c.JSON(http.StatusOK, provider)
					}
				}
			}
		}
	}
}
