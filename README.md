# satplan-server
Satplan-server is a backend application written in Go, provide APIs for mission planning.

## Database
Satplan-server share databases with [calpath](https://github.com/figwh/calpath)

## APIs
API|type|description
----|----|----
/sattree|GET|get all the satellites and sensors
/satplan|POST|misson planning
/sat/all|GET|get all satellites
/sat/add|POST|create new satellite
/sat/:id|GET|get satellite by id
/sat/update/:id|PUT|update satellite
/sat/:id|DELETE|delete satellite
/sat/tle/update|PUT|update tles
/sat/tle/cal|POST|recalculation data
/sen/all|GET|get all sensors
/sen/add|GET|get all sensors
/sen/bysat|GET|get sensors by satellite
/sen/:id|GET|get sensor by id
/sen/update/:id|PUT|update sensor
/sen/:id|DELETE|delete sensor

For more details, please take a look at syscfg/router.go

## Auto update TLEs
Satplan-server will update TLEs and recalculate data at 00:00 am UTC

## Build and run
### Prerequisites
Make sure you have installed all of the following prerequisites on your development machine:
* Go 1.15+

Run commands below to build and run
```bash
CGO_ENABLED=1 go build 
./satplan-server
```