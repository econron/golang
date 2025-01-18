curl -X POST http://localhost:8000/send \
-H "Content-Type: application/json" \
-d '{"server": 1, "order": 1}'

curl -X POST http://localhost:8000/register/server \
-H "Content-Type: application/json" \
-d '{"id": 1, "server_url": "http://localhost:8001"}'
