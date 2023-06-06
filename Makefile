build:
	cd backend && GOOS=linux GOARCH=amd64 make build

run: build
	sudo docker compose up --force-recreate --build

check:
	cd checkout && make precommit
