export green=`tput setaf 2`
export reset=`tput sgr0`

all:
	go build -o reciever main.go
	@echo "${green}reciever successfully built${reset}"
	go build -o publisher publisher.go
	@echo "${green}publisher successfully built${reset}"
	sudo docker-compose -f docker-compose.yaml up --build
up:
	docker-compose -f docker-compose.yaml up
down:
	docker-compose -f docker-compose.yaml down
clean:
	rm -rf /home/kseed/docker/db/*
	rm -rf publisher reciever
	docker volume rm db
fclean: clean
	docker rmi -f $$(docker images -qa)
