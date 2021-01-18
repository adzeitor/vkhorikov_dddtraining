package domain

import (
	"regexp"
	"time"
	"unicode/utf8"
)

type Customer struct {
	Entity

	Name                 string
	Email                string
	Status               CustomerStatus
	StatusExpirationDate *time.Time
	MoneySpent           int
	PurchasedMovies      []PurchasedMovie
}

func (c Customer) IsValid() bool {
	if utf8.RuneCountInString(c.Name) > 100 {
		return false
	}

	re := regexp.MustCompile(`^.+@.+$`)
	if !re.MatchString(c.Email) {
		return false
	}

	return true
}
