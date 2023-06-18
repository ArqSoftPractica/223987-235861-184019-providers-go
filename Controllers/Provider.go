package Controllers

import (
	"223987-235861-184019-providers/Models"
	"223987-235861-184019-providers/Repository"
	"223987-235861-184019-providers/Util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToCustomError(err error) *Util.CustomError {
	if customErr, ok := err.(*Util.CustomError); ok {
		return customErr
	}
	return &Util.CustomError{
		Message: err.Error(),
	}
}

func GetProviders(c *gin.Context) {
	var providers []Models.Provider

	err := Repository.GetAllProviders(&providers)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, ToCustomError(err))
	} else {
		c.JSON(http.StatusOK, providers)
	}
}

func CreateProvider(c *gin.Context) {
	var provider Models.Provider
	c.BindJSON(&provider)
	err := Repository.CreateProvider(&provider)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
	} else {
		c.JSON(http.StatusOK, provider)
	}
}

func GetProviderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var provider Models.Provider
	err := Repository.GetProviderByID(&provider, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, ToCustomError(err))
	} else {
		c.JSON(http.StatusOK, provider)
	}
}

func UpdateProvider(c *gin.Context) {
	var provider Models.Provider
	id := c.Params.ByName("id")
	err := Repository.GetProviderByID(&provider, id)
	if err != nil {
		c.JSON(http.StatusNotFound, provider)
	}
	c.BindJSON(&provider)
	err = Repository.UpdateProvider(&provider, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
	} else {
		c.JSON(http.StatusOK, provider)
	}
}

func DeleteProvider(c *gin.Context) {
	var provider Models.Provider
	id := c.Params.ByName("id")
	err := Repository.GetProviderByID(&provider, id)
	if err != nil {
		c.JSON(http.StatusNotFound, provider)
	}

	provider.IsActive = false

	err = Repository.UpdateProvider(&provider, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ToCustomError(err))
	} else {
		c.JSON(http.StatusOK, provider)
	}
}
