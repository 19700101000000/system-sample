dep:
	cd server/ && dep init

build:
	sudo docker-compose up -d --build

init: dep build

up:
	sudo docker-compose up -d

stop:
	sudo docker-compose stop

ps:
	sudo docker-compose ps

rm:
	sudo docker-compose rm

logs:
	suto docker-compose logs

stopup: stop up
