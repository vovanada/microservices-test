## How to run

* Create mongodb volume (once) - `docker volume create --name=port_mongo_data`
* Run docker-compose - `docker-compose up`
* Api available by url - http://localhost:8888

## Linter

* Install https://github.com/golangci/golangci-lint
* Run linter from root directory `make run-linter`

## Tests

* Run all services
* Run tests `make run-tests`

## Swagger documentation

* After run docker-compose open `http://localhost:8889` with swagger documentation

## Possible improvements

* Better error handling in gateways
* Run tests on separate database
* Add in test checking data in database after import


