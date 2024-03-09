package controller

import (
	"hacktiv-assignment-2/model"
	"hacktiv-assignment-2/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	Router       *gin.Engine
	orderUsecase usecase.OrderUsecase
}

func (c *OrderController) Create(ctx *gin.Context) {
	var order model.Order
	err := ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.orderUsecase.Create(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Order created"})
}

func (c *OrderController) FindAll(ctx *gin.Context) {
	orders, err := c.orderUsecase.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (c *OrderController) FindByID(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := c.orderUsecase.FindByID(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func (c *OrderController) Update(ctx *gin.Context) {
	var order model.Order
	orderID := ctx.Param("orderID")

	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = c.orderUsecase.FindByID(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order.OrderId = orderId
	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.orderUsecase.Update(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}

func (c *OrderController) Delete(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	orderId, err := strconv.Atoi(orderID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = c.orderUsecase.FindByID(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.orderUsecase.Delete(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted"})
}

func NewOrderController(router *gin.Engine, usecase usecase.OrderUsecase) *OrderController {
	controller := &OrderController{
		Router:       router,
		orderUsecase: usecase,
	}

	router.POST("/orders", controller.Create)
	router.GET("/orders", controller.FindAll)
	router.GET("/orders/:orderID", controller.FindByID)
	router.PUT("/orders/:orderID", controller.Update)
	router.DELETE("/orders/:orderID", controller.Delete)

	return controller
}
