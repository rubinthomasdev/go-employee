# go-employee
Clean architecture in Go.  

Based on the learnings from the book [Get Your Hands Dirty on Clean Architecture](https://www.packtpub.com/product/get-your-hands-dirty-on-clean-architecture/9781839211966) by [Tom Hombergs](https://twitter.com/TomHombergs)

# testing
## requires docker and make
```
unit tests with code coverage:
TEST_PG_HOST=localhost TEST_PG_PORT=5432 TEST_PG_USER=postgres TEST_PG_PASSWORD=testpwd TEST_PG_DBNAME=postgres make test

integration tests with code coverage:
TEST_PG_HOST=localhost TEST_PG_PORT=5432 TEST_PG_USER=postgres TEST_PG_PASSWORD=testpwd TEST_PG_DBNAME=postgres make test-integration
```


# run
## requires docker and postgres

```
docker build . -t rubinthomasdev/go-employee:latest

for postgres running on windows host
docker run -p 8080:8080 -e PG_HOST=host.docker.internal -e PG_PORT=5432 -e PG_USER=postgres -e PG_PASSWORD=yourpgpassword -e PG_DBNAME=postgres --name employeeapi rubinthomasdev/go-employee:latest

for postgres running on mac host
docker run -p 8080:8080 -e PG_HOST=docker.for.mac.localhost -e PG_PORT=5432 -e PG_USER=postgres -e PG_PASSWORD=yourpgpassword -e PG_DBNAME=postgres --name employeeapi rubinthomasdev/go-employee:latest

```

# endpoints:
### to get all employees data:
```
curl --location --request GET 'localhost:8080/api/v1/employees' \
--header 'Content-Type: application/json'
```

### to get a single employees data:
```
curl --location --request GET 'localhost:8080/api/v1/employees/1' \
--header 'Content-Type: application/json'
```
