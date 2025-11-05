package validation

import (
	"errors"
	"regexp"
	"strings"

	models "user-service/internal/model"
)

// ValidateUser validates fields of model.User
func ValidateUser(u *models.User) error {
	// first_name
	u.FirstName = strings.TrimSpace(u.FirstName)
	if len(u.FirstName) == 0 {
		return errors.New("first_name is required")
	}

	// last_name
	u.LastName = strings.TrimSpace(u.LastName)
	if len(u.LastName) == 0 {
		return errors.New("last_name is required")
	}

	// phone
	u.Phone = strings.TrimSpace(u.Phone)
	phoneRegex := regexp.MustCompile(`^\+?[0-9]{9,15}$`)
	if len(u.Phone) == 0 {
		return errors.New("phone is required")
	}
	if !phoneRegex.MatchString(u.Phone) {
		return errors.New("invalid phone format")
	}

	// email
	u.Email = strings.TrimSpace(u.Email)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if len(u.Email) == 0 {
		return errors.New("email is required")
	}
	if !emailRegex.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	// tg_username (optional, @ belgisini olib tashlab tekshiradi)
	if len(u.TgUsername) > 0 {
		tgUsername := strings.TrimPrefix(strings.TrimSpace(u.TgUsername), "@")
		tgRegex := regexp.MustCompile(`^[a-zA-Z0-9_]{5,32}$`)
		if !tgRegex.MatchString(tgUsername) {
			return errors.New("invalid tg_username format")
		}
	}

	return nil
}
