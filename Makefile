# POSTMAN_API_KEY=$(shell cat .env)

run-e2e:
	curl -H 'X-Api-Key:$(POSTMAN_API_KEY)' https://api.getpostman.com/collections/625522-4e36bf20-943e-436b-a11f-9931275992ee > e2e/api.json
	curl -H 'X-Api-Key:$(POSTMAN_API_KEY)' https://api.getpostman.com/environments/625522-5469f08b-97de-4721-99c6-fbc21c782338 > e2e/e2e.postman_environment.json
	docker-compose run newman newman run e2e/api.json --environment=e2e/e2e.postman_environment.json