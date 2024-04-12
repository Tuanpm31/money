package api

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

func (h *Handler) GetMoney(c *gin.Context) {
	id := c.Param("id")
	var amount float64
	err := h.DB.QueryRow("SELECT amount FROM money WHERE id = $1", id).Scan(&amount)
	if errors.Is(err, sql.ErrNoRows) {
		// If not found, create a new record with amount set to 0
		_, err := h.DB.Exec("INSERT INTO money (id, amount) VALUES ($1, $2)", id, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		amount = 0
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"amount": amount})
}

func (h *Handler) SetMoney(c *gin.Context) {
	id := c.Param("id")
	var payload struct {
		Amount float64 `json:"amount"`
	}
	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.DB.Exec("UPDATE money SET amount = $2 WHERE id = $1", id, payload.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"new_amount": payload.Amount})
}
