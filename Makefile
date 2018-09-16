dep:
	cd server/ && \
	dep init

build:
	sudo docker-compose up -d --build

init: dep build

ps:
	sudo docker-compose ps

up:
	sudo docker-compose up -d

stop:
	sudo docker-compose stop

rm:
	sudo docker-compose rm

logs:
	sudo docker-compose logs

stopup: stop up
