dep:
	dep ensure -v

gen:
	(cd proto; ./gen.sh)

down:
	docker-compose down

up:
	docker-compose up -d
