package trait

import (
	"errors"
	"fmt"
	"hastane-takip/internal/models"
	"os"
	"time"

	"github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IPasswordResetHandler interface {
	RequestResetCode(phone string) (string, error)
	VerifyResetCode(phone, code string) error
	ResetPassword(phone, newPassword string) error
	SendSMS(to, message string) error
}

type passwordResetHandler struct {
	DB *gorm.DB
}

func NewPasswordResetHandler(db *gorm.DB) IPasswordResetHandler {
	return &passwordResetHandler{DB: db}
}

func (prt *passwordResetHandler) RequestResetCode(phone string) (string, error) {
	resetCode := GenerateResetCode()
	expiresAt := time.Now().Add(15 * time.Minute)

	passwordReset := models.PasswordReset{
		Phone:     phone,
		Code:      resetCode,
		ExpiresAt: expiresAt,
	}

	if err := prt.DB.Where("phone = ?", phone).Delete(&models.PasswordReset{}).Error; err != nil {
		return "", err
	}

	if err := prt.DB.Create(&passwordReset).Error; err != nil {
		return "", err
	}

	message := "Your password reset code is: " + resetCode
	if err := prt.SendSMS(phone, message); err != nil {
		return "", err
	}

	return resetCode, nil
}

func (prt *passwordResetHandler) VerifyResetCode(phone, code string) error {
	var passwordReset models.PasswordReset
	if err := prt.DB.Where("phone = ? AND code = ?", phone, code).First(&passwordReset).Error; err != nil {
		return err
	}

	if time.Now().After(passwordReset.ExpiresAt) {
		return errors.New("reset code has expired")
	}

	return nil
}

func (prt *passwordResetHandler) ResetPassword(phone, newPassword string) error {
	hashedPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := prt.DB.Model(&models.Staff{}).Where("phone_number = ?", phone).Update("password", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}

func (prt *passwordResetHandler) SendSMS(to, message string) error {
	from := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetBody(message)
	params.SetFrom(from)
	params.SetTo("+15558675310")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return nil
}

func GenerateResetCode() string {
	return fmt.Sprintf("%06d", 123456)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
