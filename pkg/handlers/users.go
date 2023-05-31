package handlers

import (
	"errors"
	"github.com/Amir122002/bank/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (h *handler) GetAllUsers(c *gin.Context) {
	var users []models.User
	if err := h.DB.Find(&users).Error; err != nil {
		log.Println("getting posts from DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *handler) GetAllUser(c *gin.Context) {
	var users []models.User
	err := h.DB.Preload("Cell").Find(&users).Error
	if err != nil {
		log.Println("getting posts from DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *handler) GetUserByLogin(c *gin.Context) {
	login := c.Param("login")

	var user models.User
	err := h.DB.Preload("Cell").Where("login = ?", login).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
			return
		}

		log.Println("getting user from DB:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *handler) ReplenishUserMoney(c *gin.Context) {
	login := c.Param("login")
	replenishStr := c.Param("replenish")

	replenish, err := strconv.Atoi(replenishStr)
	if err != nil {
		log.Println("failed to convert replenish amount:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid replenish amount",
		})
		return
	}

	var user models.User
	if err := h.DB.Preload("Cell").First(&user, "login = ?", login).Error; err != nil {
		log.Println("failed to find user:", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if len(user.Cell) == 0 {
		log.Println("user has no cell")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User has no cell",
		})
		return
	}

	user.Cell[0].Money += replenish
	if err := h.DB.Save(&user.Cell[0]).Error; err != nil {
		log.Println("failed to update cell money:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update cell money",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User's account replenished",
	})
}

func (h *handler) WithdrawUserMoney(c *gin.Context) {
	login := c.Param("login")
	withdrawStr := c.Param("withdraw")

	withdraw, err := strconv.Atoi(withdrawStr)
	if err != nil {
		log.Println("failed to convert withdraw amount:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid withdraw amount",
		})
		return
	}

	var user models.User
	if err := h.DB.Preload("Cell").First(&user, "login = ?", login).Error; err != nil {
		log.Println("failed to find user:", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if len(user.Cell) == 0 {
		log.Println("user has no cell")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "User has no cell",
		})
		return
	}

	if user.Cell[0].Money < withdraw {
		log.Println("insufficient funds")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insufficient funds",
		})
		return
	}

	user.Cell[0].Money -= withdraw
	if err := h.DB.Save(&user.Cell[0]).Error; err != nil {
		log.Println("failed to update cell money:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update cell money",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User's account withdrawed",
	})
}
