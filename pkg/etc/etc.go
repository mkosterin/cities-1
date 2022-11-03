package etc

import (
	"bytes"
	"cities-1/pkg/city"
	"encoding/json"
)

func StructToBytes(s any) []byte {
	//Any struct to []bytes
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	return reqBodyBytes.Bytes()
}

func CityToCityWithId(id int, c city.City) (resp city.CityWithId) {
	resp.Id = id
	resp.Name = c.Name
	resp.Region = c.Region
	resp.District = c.District
	resp.Foundation = c.Foundation
	resp.Population = c.Population
	return
}

func CityWithIdToCity(c city.CityWithId) (resp city.City) {
	resp.Name = c.Name
	resp.Region = c.Region
	resp.District = c.District
	resp.Foundation = c.Foundation
	resp.Population = c.Population
	return
}
