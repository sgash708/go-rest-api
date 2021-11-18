build:
	docker-compose build
up:
	docker-compose up -d
down:
	docker-compose down
ls:
	docker ps -a
exec:
	docker-compose exec golang sh
logs:
	docker logs golang