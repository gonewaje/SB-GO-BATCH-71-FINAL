package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"gonewaje/final/repository"
	"gonewaje/final/structs"

	"github.com/gin-gonic/gin"
)

type MenusController struct {
	DB *sql.DB
}

func (mc *MenusController) ListByRestaurant(c *gin.Context) {
	// rid, _ := strconv.Atoi(c.Param("restaurant_id"))
	rid, _ := strconv.Atoi(c.Param("id"))
	items, err := repository.ListMenuItemsByRestaurant(mc.DB, rid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func (mc *MenusController) Create(c *gin.Context) {
	var in structs.MenuItem
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := repository.CreateMenuItem(mc.DB, in.RestaurantID, in.Name, in.Price, in.Available)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	in.ID = id
	c.JSON(http.StatusCreated, gin.H{"message": "created", "data": in})
}

func (mc *MenusController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in structs.MenuItem
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ok, err := repository.UpdateMenuItem(mc.DB, id, in.Name, in.Price, in.Available)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "menu item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (mc *MenusController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	ok, err := repository.DeleteMenuItem(mc.DB, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "menu item not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
