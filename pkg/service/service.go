package service

import (
	"golang-cities-information-service/pkg/model"
)

type Service interface {
	Create(request *model.CityRequest) (int, error)
	Delete(id int) error
	GetFull(id int) (*model.CityStruct, error)
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
}
type CityService struct {
	repo Storage
}

func NewCityService(repo Storage) *CityService {
	return &CityService{repo: repo}
}

type Storage interface {
	Create(city *model.CityRequest) (int, error)
	Delete(id int) error
	GetFull(id int) (*model.CityStruct, error)
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
	//	findCities(idx_type int, searchText string) []string

}

func (сs *CityService) Create(city *model.CityRequest) (int, error) {
	return сs.repo.Create(city)
}
func (сs *CityService) Delete(id int) error {
	return сs.repo.Delete(id)
}

func (сs *CityService) GetFull(id int) (*model.CityStruct, error) {
	return сs.repo.GetFull(id)
}

func (cs *CityService) SetPopulation(id, population int) error {
	return cs.repo.SetPopulation(id, population)
}

func (сs *CityService) GetFromRegion(region string) ([]string, error) {
	return сs.repo.GetFromRegion(region)
}

func (сs *CityService) GetFromDistrict(district string) ([]string, error) {
	return сs.repo.GetFromDistrict(district)
}

func (сs *CityService) GetFromPopulation(start, end int) ([]string, error) {
	return сs.repo.GetFromPopulation(start, end)
}

func (сs *CityService) GetFromFoundation(start, end int) ([]string, error) {
	return сs.repo.GetFromFoundation(start, end)
}
