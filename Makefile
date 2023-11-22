include .env
export $(shell sed 's/=.*//' .env)

db:
	docker run -p 3306:3306 --name graphql-tutorial-mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=graphqltutorial -d mysql:latest

db-console:
	# mysql -u root -p graphqltutorial
	docker exec -it graphql-tutorial-mysql bash 

run:
	go run ./cmd/server.go

generate:
	go get github.com/99designs/gqlgen@v0.17.40 && go run github.com/99designs/gqlgen generate

token:
	curl --request POST \
  --url 'https://${AUTH0_DOMAIN}/oauth/token' \
  --header 'content-type: application/x-www-form-urlencoded' \
  --data grant_type=client_credentials \
  --data client_id=${AUTH0_CLIENT_ID} \
  --data client_secret=${AUTH0_CLIENT_SECRET_ID} \
  --data audience=${AUTH0_AUDIENCE}

migration:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)