package controller

import (
	"OnionPractice/app/usecase"
	"OnionPractice/di"
	"OnionPractice/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCategoriesResponse struct {
	ID        int    `json:"id"`
	CompanyID int    `json:"companyID"`
	Name      string `json:"name"`
}

func GetAllCategories(c *gin.Context) {
	u, err := di.GetAllCategoriesUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))
	}

	categories, err := u.GetAllCategoriesUseCase(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))
	}

	response := make([]GetCategoriesResponse, len(categories))
	for i, category := range categories {
		response[i] = GetCategoriesResponse{
			ID:        category.ID(),
			CompanyID: category.ID(),
			Name:      category.Name(),
		}
	}

	c.JSON(http.StatusOK, response)
}

type CreateCategoryRequest struct {
	ID        int    `json:"id"`
	CompanyID int    `json:"companyID"`
	Name      string `json:"name"`
}

func CreateCategory(c *gin.Context) {
	req := CreateCategoryRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, helpers.RenderError(http.StatusBadRequest))

		return
	}

	u, err := di.CreateCategoryUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))

		return
	}

	input := usecase.CreateCategoryUseCaseInput{
		ID:        req.ID,
		CompanyID: req.CompanyID,
		Name:      req.Name,
	}

	result, err := u.CreateCategoryUseCase(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.RenderError(http.StatusInternalServerError))

		return
	}

	c.JSON(201, gin.H{
		"id":        result.Category.ID(),
		"companyID": result.Category.CompanyID(),
		"name":      result.Category.Name(),
	})
}
