package domain

import "time"

type CustomerService struct {
	MovieService MovieService
}

func NewCustomerService() CustomerService {
	return CustomerService{
		MovieService: NewMovieService(),
	}
}

func (s CustomerService) CalculatePrice(
	status CustomerStatus,
	statusExpirationDate *time.Time,
	licensingModel LicensingModel,
) int {
	var price int
	switch licensingModel {
	case LicensingModelTwoDays:
		price = 400 // $4.00
	case LicensingModelLifeLong:
		price = 800 // $8.00
	default:
		panic("unknown licensing model")
	}

	if status == CustomerStatusAdvanced &&
		(statusExpirationDate == nil || statusExpirationDate.After(time.Now())) {
		price = price * 3 / 4 // FIXME: *0.75
	}

	return price
}

func (s CustomerService) PurchaseMovie(customer *Customer, movie *Movie) {
	expirationDate := s.MovieService.GetExpirationDate(movie.LicensingModel)
	price := s.CalculatePrice(
		customer.Status, customer.StatusExpirationDate, movie.LicensingModel,
	)
	purchasedMovie := PurchasedMovie{
		MovieId:        movie.Id,
		CustomerId:     customer.Id,
		ExpirationDate: expirationDate,
		PurchaseDate:   time.Now(),
		Price:          price,
	}
	customer.PurchasedMovies = append(customer.PurchasedMovies, purchasedMovie)
	customer.MoneySpent += price
}

func (s CustomerService) purchasedMoviesCountLastDays(
	customer *Customer, days int,
) int {
	count := 0
	afterDate := time.Now().AddDate(0, 0, -days)
	for _, movie := range customer.PurchasedMovies {
		if movie.PurchaseDate.After(afterDate) {
			count += 1
		}
	}
	return count
}

func (s CustomerService) lastYearSpent(customer *Customer) int {
	sum := 0
	startOfLastYear := time.Now().AddDate(-1, 0, 0)
	for _, movie := range customer.PurchasedMovies {
		if movie.PurchaseDate.After(startOfLastYear) {
			sum += movie.Price
		}
	}
	return sum
}

func (s CustomerService) PromoteCustomer(customer *Customer) bool {
	// at least 2 active movies during last 30 days
	if s.purchasedMoviesCountLastDays(customer, 30) < 2 {
		return false
	}

	// at least 100 dollars spent during the last year
	hundredDollars := 10000 // $100.00
	if s.lastYearSpent(customer) < hundredDollars {
		return false
	}

	customer.Status = CustomerStatusAdvanced
	statusExpiration := time.Now().AddDate(1, 0, 0)
	customer.StatusExpirationDate = &statusExpiration
	return true
}
