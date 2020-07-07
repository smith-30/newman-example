run-e2e:
	curl -H 'X-Api-Key:' https://api.getpostman.com/collections/ > e2e/origin.json
	sed -e 's/"\([^"]*\)techtouch-player.zip"/"\/etc\/newman\/e2e\/techtouch-player.zip"/g' \
	    -e 's/"\([^"]*\)techtouch-editor.zip"/"\/etc\/newman\/e2e\/techtouch-editor.zip"/g' e2e/origin.json > e2e/api.json
	curl -H 'X-Api-Key:P' https://api.getpostman.com/environments/ > e2e/e2e.postman_environment.json
	docker-compose -f docker-compose_e2e.yml run newman newman run e2e/api.json --environment=e2e/e2e.postman_environment.json --folder AuthZ