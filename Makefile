
# docker build
build: 
	docker compose build

# test 
test: build run
	sleep 5 && go test -coverprofile=cover.out ./...

# integration
test-integration: build run
	sleep 5 && go test -coverprofile=cover.out -coverpkg ./.../persistence,./.../web ./integrationtests/adapter/...

# build and run 
run : build
	docker compose up -d --remove-orphans