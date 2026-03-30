#!/bin/sh

curl -f -i -X GET 'localhost:8080/api/models'

curl -f -i -X POST 'localhost:8080/api/models' --data '{"name": "test", "repository": "github.com/2Delight/example-service", "version": "v0.1.0"}'

curl -f -i -X POST 'localhost:8080/api/models' --data '{"name": "test", "repository": "github.com/2Delight/example-service", "version": "v0.1.1"}'

curl -f -i -X PUT 'localhost:8080/api/models?id=2' --data '{"name": "test", "repository": "github.com/2Delight/example-service", "version": "v0.1.2"}'

curl -f -i -X DELETE 'localhost:8080/api/models?id=1'

curl -f -i -X GET 'localhost:8080/api/models/lookup?repository=github.com/2Delight/example-service&version=v0.1.0'
