package application

import (
	"time"

	"dddtraining/domain"
)

type CustomerRepository struct {
	customers map[int]*domain.Customer
	lastId    int
}

func NewCustomerRepository() *CustomerRepository {
	repo := CustomerRepository{
		customers: make(map[int]*domain.Customer),
		lastId:    0,
	}

	twoDays := time.Now().AddDate(0, 0, 2)
	alice := domain.Customer{
		Name:   "Alice Alison",
		Email:  "alice@gmail.com",
		Status: domain.CustomerStatusRegular,
		PurchasedMovies: []domain.PurchasedMovie{
			{
				CustomerId:     1,
				ExpirationDate: &twoDays,
				MovieId:        1,
				Price:          400, // $4.00
				PurchaseDate:   time.Now(),
			},
		},
	}

	bob := domain.Customer{
		Name:   "Bob Bobson",
		Email:  "bob@gmail.com",
		Status: domain.CustomerStatusRegular,
	}

	repo.Save(&alice)
	repo.Save(&bob)
	return &repo
}

func (repo *CustomerRepository) find(
	match func(*domain.Customer) bool,
) *domain.Customer {
	for _, customer := range repo.customers {
		if match(customer) {
			return customer
		}
	}
	return nil
}

func (repo *CustomerRepository) GetById(id int) *domain.Customer {
	return repo.find(func(customer *domain.Customer) bool {
		return customer.Id == id
	})
}

func (repo *CustomerRepository) GetByEmail(email string) *domain.Customer {
	return repo.find(func(customer *domain.Customer) bool {
		return customer.Email == email
	})
}

func (repo *CustomerRepository) Save(customer *domain.Customer) {
	if customer.Id == 0 {
		repo.lastId++
		customer.SetId(repo.lastId)
	}
	repo.customers[customer.Id] = customer
}
