curl "http://localhost:8080/function/Call" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{
  "name": "my-first-func",
  "request": {}
}'