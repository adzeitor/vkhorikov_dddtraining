package domain

import "time"

type PurchasedMovie struct {
	MovieId        int
	CustomerId     int
	Price          int
	PurchaseDate   time.Time
	ExpirationDate *time.Time
}
