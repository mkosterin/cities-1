package store

import (
	"cities-1/pkg/city"
	"cities-1/pkg/etc"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"sync"
)

type cities map[int]city.City

type CityStruct struct {
	mu     sync.Mutex
	cities cities
}

func NewStore() *CityStruct {
	//Storage constructor
	var store CityStruct
	store.cities = make(cities)
	return &store
}

func (c *CityStruct) LoadFromCsv(filePath string) error {
	//read dataset from CSV
	c.mu.Lock()
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
		(*c).cities[cityId] = bufferCity
	}
	c.mu.Unlock()
	return nil
}

func (c *CityStruct) SaveToCsv(filePath string) error {
	//write store to csv
	c.mu.Lock()
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
		return err
	}
	w := csv.NewWriter(file)
	defer w.Flush()
	var data [][]string
	for key, record := range (*c).cities {
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
	c.mu.Unlock()
	return nil
}

func (c *CityStruct) IsExistsById(id int) (resp bool) {
	//If city exists == true, otherwise == false
	c.mu.Lock()
	_, resp = (*c).cities[id]
	c.mu.Unlock()
	return
}

func (c *CityStruct) GetById(id int) (resp city.CityWithId) {
	//return CityWithId struct
	c.mu.Lock()
	resp = etc.CityToCityWithId(id, (*c).cities[id])
	c.mu.Unlock()
	return
}

func (c *CityStruct) NewById(city *city.CityWithId) {
	//create new record in database
	id := city.Id
	(*c).cities[id] = etc.CityWithIdToCity(*city)
}

func (c *CityStruct) DeleteById(id int) {
	//delete record from database by id
	c.mu.Lock()
	delete((*c).cities, id)
	c.mu.Unlock()
}

func (c *CityStruct) UpdatePopulationById(id, population int) {
	//Update population by id
	c.mu.Lock()
	entry := (*c).cities[id]
	entry.Population = population
	(*c).cities[id] = entry
	c.mu.Unlock()
}

func (c *CityStruct) GetCitiesFromRegion(region string) (resp []string) {
	//Return list of cities from region
	c.mu.Lock()
	for _, value := range (*c).cities {
		if value.Region == region {
			resp = append(resp, value.Name)
		}
	}
	c.mu.Unlock()
	return
}

func (c *CityStruct) GetCitiesFromDistrict(district string) (resp []string) {
	//Return list of cities from district
	c.mu.Lock()
	for _, value := range (*c).cities {
		if value.District == district {
			resp = append(resp, value.Name)
		}
	}
	c.mu.Unlock()
	return
}

func (c *CityStruct) GetCitiesByPopulation(min, max int) (resp []string) {
	//Return list of cities with population in the range
	c.mu.Lock()
	for _, value := range (*c).cities {
		if value.Population >= min && value.Population <= max {
			resp = append(resp, value.Name)
		}
	}
	c.mu.Unlock()
	return
}

func (c *CityStruct) GetCitiesByFoundation(min, max int) (resp []string) {
	//Return list of cities with foundation in the range
	c.mu.Lock()
	for _, value := range (*c).cities {
		if value.Foundation >= min && value.Foundation <= max {
			resp = append(resp, value.Name)
		}
	}
	c.mu.Unlock()
	return
}
