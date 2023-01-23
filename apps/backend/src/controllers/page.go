package controllers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageInput struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// GetPage godoc
// @Summary 		List pages
// @Description get pages
// @Tags 				pages
// @Accept 			json
// @Produce 		json
// @Success     200  {array}   models.Page
// @Router 			/pages [get]
func GetPage(c *gin.Context) {
	var pages []models.Page
	models.DB.Find(&pages)

	c.JSON(http.StatusOK, pages)
}

// PostPage godoc
// @Summary 		Create an page
// @Description post an page
// @Tags 				pages
// @Accept 			json
// @Produce 		json
// @Param      	body   body   PageInput  true  "Page"
// @Success 		201  {int}  http.StatusCreated
// @Failure     400  {int}  http.StatusBadRequest
// @Router 			/pages [post]
func PostPage(c *gin.Context) {
	// Validate input
	var input PageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add the new page
	page := models.Page{Title: input.Title, Body: input.Body}
	models.DB.Create(&page)

	c.JSON(http.StatusCreated, page)
}

// GetPageByTitle godoc
// @Summary 		Show an page
// @Description get page by Title
// @Tags 				pages
// @Accept 			json
// @Produce 		json
// @Param      	title   path   string  true  "Page title"
// @Success     200  {object}  models.Page
// @Failure     404  {int}  http.StatusNotFound
// @Router 			/pages/{title} [get]
func GetPageByTitle(c *gin.Context) {
	var page models.Page

	if err := models.DB.Where("title = ?", c.Param("title")).First(&page).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found!"})
		return
	}

	c.JSON(http.StatusOK, page)
}

// DeletePage godoc
// @Summary 		Delete an page
// @Description delete an page
// @Tags 				pages
// @Accept 			json
// @Produce 		json
// @Param      	title   path      string  true  "Page Title"
// @Success     200  {object}  models.Page
// @Failure     400  {int}  http.StatusBadRequest
// @Router 			/pages/{title} [delete]
func DeletePage(c *gin.Context) {
	var page models.Page

	if err := models.DB.Where("title = ?", c.Param("title")).First(&page).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page not found!"})
		return
	}

	models.DB.Delete(&page)

	c.JSON(http.StatusOK, gin.H{"message": "page deleted"})
}

// PatchPage godoc
// @Summary 		Update an page
// @Description update an page
// @Tags 				pages
// @Accept 			json
// @Produce 		json
// @Param      	title     path      string  true  "Page ID"
// @Param      	body   body   PageInput  true  "Page"
// @Success     200  {object}  models.Page
// @Failure     500  {int}  http.StatusBadRequest
// @Router 			/pages/{title} [patch]
func PatchPage(c *gin.Context) {
	var page models.Page

	if err := models.DB.Where("title = ?", c.Param("title")).First(&page).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&page).Updates(input)

	c.JSON(http.StatusOK, page)
}

// PutPage godoc
// @Summary 		Update an page
// @Description update an page
// @Tags 				pages
// @Accept 			json
// @Produce 		json
// @Param      	title   path      string  true  "Page Title"
// @Param      	body body   PageInput  true  "Page"
// @Success     200  {object}  models.Page
// @Failure     400  {int}  http.StatusBadRequest
// @Router 			/pages/{title} [put]
func PutPage(c *gin.Context) {
	var page models.Page

	if err := models.DB.Where("title = ?", c.Param("title")).First(&page).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input PageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&page).Updates(input)

	c.JSON(http.StatusOK, page)
}
