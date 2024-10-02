# Sample flow of API calls

1. curl -X POST http://localhost:8080/register -H 'Content-Type: application/json' -d '{
  "username": "user1",                                                                                                                            "password": "password1",
  "role": "write"
}'
{"message":"User registered successfully"}

2. curl -X POST http://localhost:8080/login -H 'Content-Type: application/json' -d '{
  "username": "user1",
  "password": "password1"
}'
{"token":"Token123"}

3. curl -X POST http://localhost:8080/posts -H 'Content-Type: application/json' -H 'Authorization: 'Token123' -d '{
  "title": "k8s operators",
  "content": "This post is about k8s operators",
  "author": "user1"
}'
{"id":"452016e9-084a-4c90-a50a-f4d38961db17","title":"k8s operators","content":"This post is about k8s operators","author":"user1","created_at":"2024-10-02T10:01:03.513570361+05:30","modified_at":"2024-10-02T10:01:03.513570361+05:30"}

4. curl -X PUT http://localhost:8080/posts/{e9bdb74b-5874-4dae-8303-7bed819fb62d} -H 'Content-Type: application/json' -d '{
  "title": "k8s operators",
  "content": "Updated content for k8s-operators",
  "author": "Anchal"
}'
{"id":"e9bdb74b-5874-4dae-8303-7bed819fb62d","title":"k8s operators","content":"Updated content for k8s-operators","author":"Anchal","created_at":"0001-01-01T00:00:00Z","modified_at":"2024-10-02T10:08:20.494949074+05:30"}

5. curl -X GET http://localhost:8080/posts -H 'Authorization: Token123'
[{"id":"e9bdb74b-5874-4dae-8303-7bed819fb62d","title":"k8s operators","content":"Updated content for k8s-operators","author":"Anchal","created_at":"0001-01-01T00:00:00Z","modified_at":"2024-10-02T10:08:20.494949074+05:30"},{"id":"452016e9-084a-4c90-a50a-f4d38961db17","title":"k8s operators","content":"This post is about k8s operators","author":"user1","created_at":"2024-10-02T10:01:03.513570361+05:30","modified_at":"2024-10-02T10:01:03.513570361+05:30"}]

6.  curl -X GET http://localhost:8080/posts/452016e9-084a-4c90-a50a-f4d38961db17 -H 'Authorization: Token123'