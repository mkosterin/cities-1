package store

import (
	"cities-1/pkg/city"
	"cities-1/pkg/etc"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cities map[int]city.City

func NewStore() *Cities {
	//Storage constructor
	store := make(Cities)
	return &store
}

func (c *Cities) LoadFromCsv(filePath string) error {
	//read dataset from CSV
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		return err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
		return err
	}
	for i := 0; i < len(records); i++ {
		var bufferCity city.City
		cityId, errC := strconv.Atoi(records[i][0])
		if errC != nil {
			log.Fatal("Unable to parse CityId "+records[i][0], errC)
			return errC
		}
		bufferCity.Name = records[i][1]
		bufferCity.Region = records[i][2]
		bufferCity.District = records[i][3]
		bufferCity.Population, errC = strconv.Atoi(records[i][4])
		if errC != nil {
			log.Fatal("Unable to parse Population "+records[i][4], errC)
			return errC
		}
		bufferCity.Foundation, errC = strconv.Atoi(records[i][5])
		if errC != nil {
			log.Fatal("Unable to parse Foundation "+records[i][5], errC)
			return errC
		}
		(*c)[cityId] = bufferCity
	}
	return nil
}

func (c *Cities) AddCity(cityWithId city.CityWithId) {
	//add new city in to storage
	(*c)[cityWithId.Id] = city.City{
		Name:       cityWithId.Name,
		Region:     cityWithId.Region,
		District:   cityWithId.District,
		Population: cityWithId.Population,
		Foundation: cityWithId.Foundation,
	}
	return
}

func (c *Cities) SaveToCsv(filePath string) error {
	//write store to csv
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
		return err
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	var data [][]string
	for key, record := range *c {
		row := []string{
			strconv.Itoa(key),
			record.Name,
			record.Region,
			record.District,
			strconv.Itoa(record.Population),
			strconv.Itoa(record.Foundation),
		}
		data = append(data, row)
	}
	err = w.WriteAll(data)
	if err != nil {
		log.Fatal("Unable to write CSV file ", err)
		return err
	}
	return nil
}

func (c *Cities) IsExistsById(id int) (resp bool) {
	//If city exists == true, otherwise == false
	_, resp = (*c)[id]
	return
}

func (c *Cities) GetById(id int) (resp city.CityWithId) {
	//return CityWithId struct
	resp = etc.CityToCityWithId(id, (*c)[id])
	return
}

func (c *Cities) NewById(city *city.CityWithId) {
	//create new record in database
	id := city.Id
	(*c)[id] = etc.CityWithIdToCity(*city)
}

func (c *Cities) DeleteById(id int) {
	//delete record from database by id
	delete(*c, id)
}

func (c *Cities) UpdatePopulationById(id, population int) {
	//Update population by id
	entry := (*c)[id]
	fmt.Println(entry)
	entry.Population = population
	(*c)[id] = entry
	fmt.Println((*c)[id])
}

func (c *Cities) GetCitiesFromRegion(region string) (resp []string) {
	//Return list of cities from region
	for _, value := range *c {
		if value.Region == region {
			resp = append(resp, value.Name)
		}
	}
	return
}

func (c *Cities) GetCitiesFromDistrict(district string) (resp []string) {
	//Return list of cities from district
	for _, value := range *c {
		if value.District == district {
			resp = append(resp, value.Name)
		}
	}
	return
}

func (c *Cities) GetCitiesByPopulation(min, max int) (resp []string) {
	//Return list of cities with population in the range
	for _, value := range *c {
		if value.Population >= min && value.Population <= max {
			resp = append(resp, value.Name)
		}
	}
	return
}

func (c *Cities) GetCitiesByFoundation(min, max int) (resp []string) {
	//Return list of cities with foundation in the range
	for _, value := range *c {
		if value.Foundation >= min && value.Foundation <= max {
			resp = append(resp, value.Name)
		}
	}
	return
}
