package utils

import (
	"errors"
	models "hastane-takip/internal/models"
	trait "hastane-takip/internal/trait"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// GenerateToken generates a JWT token for a given staff ID
func GenerateToken(staffID uint, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"staff_id": staffID,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses and validates a JWT token
func ParseToken(tokenString string, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

// GetStaffFromToken extracts staff information from the JWT token
func GetStaffFromToken(c *fiber.Ctx, secretKey string, staffHandler *trait.StaffHandler) (*models.Staff, error) {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("missing Authorization header")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return nil, errors.New("invalid Authorization header format")
	}

	tokenString := parts[1]

	// Parse the JWT token
	claims, err := ParseToken(tokenString, secretKey)
	if err != nil {
		return nil, err
	}

	// Get the staff ID from claims
	staffID, ok := claims["staff_id"].(float64)
	if !ok {
		return nil, errors.New("staff ID not found in token")
	}

	// Convert staffID from float64 to uint
	id := uint(staffID)

	// Fetch the staff from the database using StaffHandler
	staff, err := staffHandler.GetStaffByID(id)
	if err != nil {
		return nil, err
	}

	return staff, nil
}

// GetStaffID retrieves the staff ID from the JWT token claims
func GetStaffID(c *fiber.Ctx) (uint, error) {
	claims, ok := c.Locals("claims").(jwt.MapClaims)
	if !ok {
		return 0, errors.New("failed to retrieve claims from token")
	}
	staffID, ok := claims["staff_id"].(float64)
	if !ok {
		return 0, errors.New("staff_id not found in token")
	}

	return uint(staffID), nil
}
