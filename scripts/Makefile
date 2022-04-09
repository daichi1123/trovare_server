filename = ../go.mod

start:
	docker-compose build
	docker-compose up -d
	docker-compose exec app go mod init
	docker-compose exec app go mod tidy