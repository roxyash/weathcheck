build:
	go build -o .bin/main cmd/main.go
run: build
	 ./.bin/main

# swag:
# 	swag init -g cmd/main.go
swag:
	swag init -g cmd/main.go


weatherchecker_build:
	docker build -t getblock .

weatherchecker:weatherchecker_build
		 docker run -dt --publish 8000:8000 --name weatherchecker weatherchecker