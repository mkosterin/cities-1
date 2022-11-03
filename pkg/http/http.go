package http

import (
	"cities-1/pkg/city"
	"cities-1/pkg/etc"
	"cities-1/pkg/store"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Router(host string, st *store.Cities) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// curl http://localhost:8080
		w.Write([]byte("use exact targets"))
	})

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		//return user from database
		//curl http://localhost:8080/cityId
		uri := strings.Split(r.RequestURI, "/")
		cityId, err := strconv.Atoi(uri[len(uri)-1])
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("id has to be digit"))
		} else if st.IsExistsById(cityId) == false {
			w.WriteHeader(404)
			w.Write([]byte("id was not found in database"))
		} else {
			cityStr := st.GetById(cityId)
			w.WriteHeader(200)
			w.Write(etc.StructToBytes(cityStr))
		}
	})

	r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
		//add new city in database
		//curl -v \
		//-d '{"id":628000, "name":"Сургут", "region":"ХМАО", "district":"Тюменский", "population":400000, "foundation":1600}' \
		//-X POST http://localhost:8080/add
		request := city.CityWithId{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Unmarshal error"))
		} else if st.IsExistsById(request.Id) {
			w.WriteHeader(405)
			w.Write([]byte("id already exists in database"))
		} else {
			w.WriteHeader(200)
			st.NewById(&request)
			log.Printf("city added to database", request)
		}
	})

	r.Delete("/*", func(w http.ResponseWriter, r *http.Request) {
		//delete city from database
		//curl -v -X DELETE http://localhost:8080/id
		request := city.CityWithId{}
		err := json.NewDecoder(r.Body).Decode(&request)
		uri := strings.Split(r.RequestURI, "/")
		cityId, err := strconv.Atoi(uri[len(uri)-1])
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("id has to be digit"))
		} else if st.IsExistsById(cityId) == false {
			w.WriteHeader(404)
			w.Write([]byte("id was not found in database"))
		} else {
			st.DeleteById(cityId)
			w.WriteHeader(200)
			log.Printf("city deleted from database", cityId)
		}
	})

	r.Put("/population", func(w http.ResponseWriter, r *http.Request) {
		//update population of the city
		//curl -v -d '{"id":628000, "population":400000}' -X PUT http://localhost:8080/population
		type req struct {
			Id         int `json:"id"`
			Population int `json:"population"`
		}
		request := req{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Unmarshal error"))
		} else if st.IsExistsById(request.Id) == false {
			w.WriteHeader(404)
			w.Write([]byte("id not found in database"))
		} else {
			w.WriteHeader(200)
			st.UpdatePopulationById(request.Id, request.Population)
			log.Printf("city population updated", request)
		}
	})

	r.Post("/citiesbyregion", func(w http.ResponseWriter, r *http.Request) {
		//get cities from region
		//curl -v -d '{"region":"ХМАО"}' -X POST http://localhost:8080/citiesbyregion
		type req struct {
			Region string `json:"region"`
		}
		request := req{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Unmarshal error"))
		} else {
			w.WriteHeader(200)
			buffer := st.GetCitiesFromRegion(request.Region)
			w.Write(etc.StructToBytes(buffer))
		}
	})

	r.Post("/citiesbydistrict", func(w http.ResponseWriter, r *http.Request) {
		//get cities from district
		//curl -v -d '{"district":"Сибирский"}' -X POST http://localhost:8080/citiesbydistrict
		type req struct {
			District string `json:"district"`
		}
		request := req{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Unmarshal error"))
		} else {
			w.WriteHeader(200)
			buffer := st.GetCitiesFromDistrict(request.District)
			w.Write(etc.StructToBytes(buffer))
		}
	})

	r.Post("/citiesbypopulation", func(w http.ResponseWriter, r *http.Request) {
		//get cities with population in range
		//curl -v -d '{"min":1000, "max":1000000}' -X POST http://localhost:8080/citiesbypopulation
		type req struct {
			Min int `json:"min"`
			Max int `json:"max"`
		}
		request := req{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Unmarshal error"))
		} else {
			w.WriteHeader(200)
			buffer := st.GetCitiesByPopulation(request.Min, request.Max)
			w.Write(etc.StructToBytes(buffer))
		}
	})

	r.Post("/citiesbyfoundation", func(w http.ResponseWriter, r *http.Request) {
		//get cities with foundation in range
		//curl -v -d '{"min":1000, "max":1000000}' -X POST http://localhost:8080/citiesbyfoundation
		type req struct {
			Min int `json:"min"`
			Max int `json:"max"`
		}
		request := req{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			w.WriteHeader(405)
			w.Write([]byte("Unmarshal error"))
		} else {
			w.WriteHeader(200)
			buffer := st.GetCitiesByFoundation(request.Min, request.Max)
			w.Write(etc.StructToBytes(buffer))
		}
	})

	http.ListenAndServe(host, r)
}
