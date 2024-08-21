

curl -XPOST http://localhost:8080/owner/profile -d `{"id":1, "name":"tekitou2","email":"tekitou@tekitou.com", "password": "password"}`


curl http://localhost:8080/user/profile \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id":1, "name":"tekitou2"}'