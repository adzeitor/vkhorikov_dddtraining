package application

import (
	"fmt"
	"time"

	"dddtraining/domain"
)

type CustomerController struct {
	MovieRepository    *MovieRepository
	CustomerRepository *CustomerRepository
	CustomerService    domain.CustomerService
}

func NewCustomerController() CustomerController {
	controller := CustomerController{
		MovieRepository:    NewMovieRepository(),
		CustomerRepository: NewCustomerRepository(),
		CustomerService:    domain.NewCustomerService(),
	}
	return controller
}

// GET /api/customers/{id}
func (c CustomerController) Get(id int) (*domain.Customer, error) {
	customer := c.CustomerRepository.GetById(id)
	if customer == nil {
		return nil, fmt.Errorf("invalid customer id: %d", id)
	}
	return customer, nil
}

// POST /api/customers
func (c CustomerController) Create(customer domain.Customer) error {
	if !customer.IsValid() {
		return fmt.Errorf("invalid model")
	}

	if c.CustomerRepository.GetByEmail(customer.Email) != nil {
		return fmt.Errorf("email is already in use: %s", customer.Email)
	}

	customer.Status = domain.CustomerStatusRegular
	c.CustomerRepository.Save(&customer)
	return nil
}

// PUT /api/customers/{id}
func (c CustomerController) Update(
	customerId int, item domain.Customer,
) error {
	if !item.IsValid() {
		return fmt.Errorf("invalid model")
	}

	customer := c.CustomerRepository.GetById(customerId)
	if customer == nil {
		return fmt.Errorf("invalid customer id: %d", customerId)
	}

	customer.Name = item.Name
	c.CustomerRepository.Save(customer)
	return nil
}

// POST /api/purchase?customerId={customerId}&movieId={movieId}
func (c CustomerController) PurchaseMovie(
	customerId int,
	movieId int,
) error {
	customer := c.CustomerRepository.GetById(customerId)
	if customer == nil {
		return fmt.Errorf("invalid customer id: %d", customerId)
	}

	movie := c.MovieRepository.GetById(movieId)
	if movie == nil {
		return fmt.Errorf("invalid movie id: %d", movieId)
	}

	alreadyPurchased := false
	for _, purchasedMovie := range customer.PurchasedMovies {
		if purchasedMovie.MovieId == movie.Id &&
			(purchasedMovie.ExpirationDate == nil ||
				purchasedMovie.ExpirationDate.After(time.Now())) {
			alreadyPurchased = true
			break
		}
	}
	if alreadyPurchased {
		return fmt.Errorf("the movie is already purchased: %s", movie.Name)
	}

	c.CustomerService.PurchaseMovie(customer, movie)

	c.CustomerRepository.Save(customer)

	return nil
}

// POST /api/promotion?customerId={customerId}
func (c CustomerController) PromoteCustomer(customerId int) error {
	customer := c.CustomerRepository.GetById(customerId)
	if customer == nil {
		return fmt.Errorf("invalid customer id: %d", customerId)
	}

	if customer.Status == domain.CustomerStatusAdvanced &&
		(customer.StatusExpirationDate == nil ||
			customer.StatusExpirationDate.Before(time.Now())) {
		return fmt.Errorf("customer already has the Advanced status")
	}

	success := c.CustomerService.PromoteCustomer(customer)
	if !success {
		return fmt.Errorf("cannot promote customer")
	}

	c.CustomerRepository.Save(customer)

	return nil
}
