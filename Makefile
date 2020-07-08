POSTMAN_API_KEY=$(shell cat .env)

run-e2e:
	curl -H 'X-Api-Key:$(POSTMAN_API_KEY)' https://api.getpostman.com/collections/9403164-841d7433-4ca3-470e-b15f-bfed26658353 > e2e/api.json
	curl -H 'X-Api-Key:$(POSTMAN_API_KEY)' https://api.getpostman.com/environments/9403164-4ed6baf7-601f-4227-a7af-9769a4cea659 > e2e/e2e.postman_environment.json
	docker-compose run newman newman run e2e/api.json --environment=e2e/e2e.postman_environment.json