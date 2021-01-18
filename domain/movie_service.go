package domain

import "time"

type MovieService struct {
}

func NewMovieService() MovieService {
	return MovieService{}
}

func (s MovieService) GetExpirationDate(model LicensingModel) *time.Time {
	twoDays := time.Now().AddDate(0, 0, 2)
	switch model {
	case LicensingModelTwoDays:
		return &twoDays
	case LicensingModelLifeLong:
		return nil
	default:
		panic("GetExpirationDate: unknown licensing")
	}
	return nil
}
