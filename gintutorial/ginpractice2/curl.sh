curl http://localhost:8080/user/profile \
    --include \
    --header "Content-Type: application/json" \
    --request "PUT" \
    --data '{"id":1, "name":"tekitou233"}'