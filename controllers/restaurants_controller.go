package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"gonewaje/final/repository"
	"gonewaje/final/structs"

	"github.com/gin-gonic/gin"
)

type RestaurantsController struct {
	DB *sql.DB
}

func (rc *RestaurantsController) List(c *gin.Context) {
	items, err := repository.ListRestaurants(rc.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (rc *RestaurantsController) Create(c *gin.Context) {
	var r structs.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := repository.CreateRestaurant(rc.DB, r.Name, r.Address, r.Phone)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	r.ID = id
	c.JSON(http.StatusCreated, gin.H{"message": "created", "data": r})
}

func (rc *RestaurantsController) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	r, err := repository.GetRestaurant(rc.DB, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": r})
}

func (rc *RestaurantsController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var r structs.Restaurant
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, err := repository.UpdateRestaurant(rc.DB, id, r.Name, r.Address, r.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (rc *RestaurantsController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ok, err := repository.DeleteRestaurant(rc.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
