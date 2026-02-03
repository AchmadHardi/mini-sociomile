package handler

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

var jwtKey = []byte("secret-key")

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var userID, tenantID int64
	var role, hashedPassword string

	err := h.db.QueryRow(`
	SELECT id, tenant_id, role, password
	FROM users
	WHERE email = ?
`, req.Email).Scan(&userID, &tenantID, &role, &hashedPassword)

	if err != nil {
		log.Println("DB ERROR:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found"})
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		return
	}

	claims := jwt.MapClaims{
		"user_id":   userID,
		"tenant_id": tenantID,
		"role":      role,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
