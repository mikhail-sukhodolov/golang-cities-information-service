package storage

import (
	"fmt"
	"github.com/pkg/errors"
	"golang-cities-information-service/pkg/model"
	//"github.com/pkg/errors"
)

type MemStorage struct {
	lastID  int
	Records map[int]*model.CityStruct
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		Records: make(map[int]*model.CityStruct),
	}
}

func (ms *MemStorage) newId() int {
	var max int
	for _, j := range ms.Records {
		if j.Id > max {
			max = j.Id
		}
	}
	fmt.Println("max", max)
	return max + 1
}

func (ms *MemStorage) Create(city *model.CityRequest) (int, error) {
	newID := ms.newId()
	NewCity := model.CityStruct{
		Id:         newID,
		Name:       city.Name,
		Region:     city.Region,
		District:   city.District,
		Population: city.Population,
		Foundation: city.Foundation,
	}
	ms.Records[len(ms.Records)+1] = &NewCity
	fmt.Println("Создан город", NewCity.Name)
	return newID, nil
}

func (ms *MemStorage) Delete(id int) error {

	for index, cityLine := range ms.Records {
		if cityLine.Id == id {
			delete(ms.Records, index)
			return nil
			break
		}
	}
	return errors.New("errNotFoundId")

}

func (ms *MemStorage) GetFull(id int) (*model.CityStruct, error) {

	for _, cityLine := range ms.Records {
		if cityLine.Id == id {
			fmt.Printf("Название города: %v\nОбласть: %v\nОкруг: %v\nНаселение: %v\nГод основания: %v\n", cityLine.Name, cityLine.Region, cityLine.District, cityLine.Population, cityLine.Foundation)
			return cityLine, nil
			break
		}
	}
	return nil, errors.New("errNotFoundId")

}

func (ms *MemStorage) SetPopulation(id, population int) error {

	for _, cityLine := range ms.Records {
		if cityLine.Id == id {
			cityLine.Population = population
			fmt.Printf("Для города %v установлено новое значение населения - %v\n", cityLine.Name, population)
			return nil
			break
		}
	}
	return errors.New("errNotFoundId")

}

func (ms *MemStorage) GetFromRegion(region string) ([]string, error) {
	cityNames := make([]string, 0)
	for _, cityLine := range ms.Records {
		if cityLine.Region == region {
			cityNames = append(cityNames, cityLine.Name)
		}
	}
	fmt.Println("Список городов по выбранному региону:", cityNames)
	return cityNames, nil
}

func (ms *MemStorage) GetFromDistrict(district string) ([]string, error) {
	cityNames := make([]string, 0)
	for _, cityLine := range ms.Records {
		if cityLine.District == district {
			cityNames = append(cityNames, cityLine.Name)
		}
	}
	fmt.Println("Список городов по выбранному региону:", cityNames)
	return cityNames, nil
}

func (ms *MemStorage) GetFromPopulation(start, end int) ([]string, error) {

	if start > end {
		start, end = end, start
	}

	cityNames := make([]string, 0)

	for _, cityLine := range ms.Records {
		if cityLine.Population >= start && cityLine.Population <= end {
			cityNames = append(cityNames, cityLine.Name)
		}
	}
	fmt.Println("Выбранному диапазону соответствуют следующие города:", cityNames)
	return cityNames, nil

}

func (ms *MemStorage) GetFromFoundation(start, end int) ([]string, error) {

	if start > end {
		start, end = end, start
	}

	cityNames := make([]string, 0)

	for _, cityLine := range ms.Records {
		if cityLine.Foundation >= start && cityLine.Foundation <= end {
			cityNames = append(cityNames, cityLine.Name)
		}
	}
	fmt.Println("Выбранному диапазону соответствуют следующие города:", cityNames)
	return cityNames, nil

}
