package application

import "dddtraining/domain"

type MovieRepository struct{}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

func (repo *MovieRepository) allMovies() []domain.Movie {
	return []domain.Movie{
		{
			Entity:         domain.Entity{Id: 1},
			Name:           "Great Gatsby",
			LicensingModel: domain.LicensingModelTwoDays,
		},
		{
			Entity:         domain.Entity{Id: 2},
			Name:           "Secret life of Pets",
			LicensingModel: domain.LicensingModelLifeLong,
		},
	}
}

func (repo *MovieRepository) find(
	match func(domain.Movie) bool,
) *domain.Movie {
	for _, movie := range repo.allMovies() {
		if match(movie) {
			return &movie
		}
	}
	return nil
}

func (repo *MovieRepository) GetById(id int) *domain.Movie {
	return repo.find(func(movie domain.Movie) bool {
		return movie.Id == id
	})
}
