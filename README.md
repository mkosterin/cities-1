## Cities REST server

### ToDo:
- load data from CSV
- ctrl-c hadler
- save data to CSV
- REST server
    * get info about city, using id
    * add new city into database
    * delete city from database
    * upgrade city population
    * get list of cities in the region
    * get list of cities in the district
    * get list of cities with population in range
    * get list of cities with foundation year in range

### HTTP Method definition:
- curl http://localhost:8080/cityId
- curl -v -d '{"id":628000, "name":"Сургут", "region":"ХМАО", "district":"Тюменский", "population":400000, "foundation":1600}' -X POST http://localhost:8080/add
- curl -v -X DELETE http://localhost:8080/id
- curl -v -d '{"id":628000, "population":400000}' -X PUT http://localhost:8080/population
- curl -v -d '{"region":"ХМАО"}' -X POST http://localhost:8080/citiesbyregion
- curl -v -d '{"district":"Сибирский"}' -X POST http://localhost:8080/citiesbydistrict
- curl -v -d '{"min":1000, "max":1000000}' -X POST http://localhost:8080/citiesbypopulation
- curl -v -d '{"min":1000, "max":1000000}' -X POST http://localhost:8080/citiesbyfoundation

### TBD:
- reduce code reducing repeatable code