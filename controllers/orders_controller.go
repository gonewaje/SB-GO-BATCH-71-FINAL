package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"gonewaje/final/repository"
	"gonewaje/final/structs"

	"github.com/gin-gonic/gin"
)

type OrdersController struct {
	DB *sql.DB
}

func (oc *OrdersController) Create(c *gin.Context) {
	var req structs.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	uid, _ := c.Get("user_id")
	orderID, total, err := repository.CreateOrderWithItems(oc.DB, uid.(int), req.RestaurantID, req.Items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "order created", "order_id": orderID, "total_price": total})
}

func (oc *OrdersController) MyOrders(c *gin.Context) {
	uid, _ := c.Get("user_id")
	orders, err := repository.ListOrdersByUser(oc.DB, uid.(int))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func (oc *OrdersController) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	if err := repository.UpdateOrderStatus(oc.DB, id, req.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status updated", "status": req.Status})
}
