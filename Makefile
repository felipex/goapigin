build:
	docker build -f ./Dockerfile -t goapigin .

run:
	docker run --rm -p 8080:8080 \
		--expose 8080 \
		--name goapigin \
		goapigin

add:
	curl -H 'Content-Type: application/json' \
		-d '{"id": 1, "user_id": "acadx0", "user_name": "aca"}' \
		-X POST \
		-v \
		http://localhost:8080/api/v1/user

get:
	curl -H 'Content-Type: application/json' \
		-X GET \
		-v \
		http://localhost:8080/api/v1/user/acadx0
