steps:

1. docker-compose build mysql
2. docker-compose up mysql
3. docker-compose build app
4. docker-compose up app

request to postman: http://localhost:8000/imports/
send following parameters as form-data:
"name": "any_value"
"email": "any_value"

wait for response,
if successful, check whether the data is inserted in 'test_db.info' table.