DB_CONTAINER_NAME = system-sample_db_1
DB_NAME = system_sample

dep:
	cd api/ && \
	dep init

build:
	sudo docker-compose up -d --build

init: dep build initdb
reset:
	cd api && \
	rm Gopkg.toml && \
	rm Gopkg.lock && \
	sudo rm -r vendor

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

restart: 
	sudo docker-compose restart

initdb:
	cat db/create_db.sql | sudo docker exec -i $(DB_CONTAINER_NAME) mysql -h 127.0.0.1 -uroot -proot
	cat db/create_table.sql | sudo docker exec -i $(DB_CONTAINER_NAME) mysql -h 127.0.0.1 -uroot -proot $(DB_NAME)
dataset:
	cat db/insert_data.sql | sudo docker exec -i $(DB_CONTAINER_NAME) mysql -h 127.0.0.1 -uroot -proot $(DB_NAME)

