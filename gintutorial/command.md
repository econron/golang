## tutorial

https://go.dev/doc/tutorial/web-service-gin

## memo

```
curl localhost:8080/albums

curl http://localhost:8080/albums --include --header "Content-Type: application/json" --request "POST" --data '{"id":"4", "title":"testing","artist":"testingartist","price":45}'

curl localhost:8080/albums/4
```