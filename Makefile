dev:
	go build -o bin/DINO dino.go
	./bin/DINO

create:
	npx sequelize db:create

drop:
	npx sequelize db:drop

migrate:
	npx sequelize db:migrate

unmigrate:
	npx sequelize db:migrate:undo:all