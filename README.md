# Go Auth Service

## How to Run (One Command)

```
docker compose up --build
```

The service will be available at http://localhost:8080

## Test the API (Curl Commands)

### 1. Sign Up
```
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"email":"konthalapalli@gmail.com", "password" : "abhishek"}'
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"email":"ajith@gmail.com", "password" : "abhishek"}'
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"email":"abhishek777@gmail.com", "password" : "abhishek"}'
```

### 2. Sign In (Get JWT and Refresh Token)
```
curl -X POST http://localhost:8080/api/auth/signin -H "Content-Type: application/json" -d '{"email":"konthalapalli@gmail.com", "password" : "abhishek"}'
```
Response will include `token` (JWT) and `refresh_token`.

### 3. Access Protected Endpoint
```
curl -H "Authorization: Bearer <JWT_TOKEN>" http://localhost:8080/api/protected
```
Replace `<JWT_TOKEN>` with the token from the signin response.

### 4. Refresh Token
```
curl -X POST -H "Content-Type: application/json" -d '{"refresh_token" : "<REFRESH_TOKEN>"}' http://localhost:8080/api/auth/refresh
```
Replace `<REFRESH_TOKEN>` with the refresh token from the signin response.

It should look something like the below: 
curl -X POST -H "Content-Type: application/json" -d '{"refresh_token" : "cZYlveu-6G5t6bqwSnwKsGvsn2Q4BsGWSyw6gfGdR5o="}' http://localhost:8080/api/auth/refresh



### 5. Revoke Token
```
curl -X POST -H "Content-Type: application/json" -d '{"refresh_token" : "cZYlveu-6G5t6bqwSnwKsGvsn2Q4BsGWSyw6gfGdR5o="}' http://localhost:8080/api/auth/revoke
```
Replace `<REFRESH_TOKEN>` with the refresh token you want to revoke.

It should look something like the below: 
curl -X POST -H "Content-Type: application/json" -d '{"refresh_token" : "cZYlveu-6G5t6bqwSnwKsGvsn2Q4BsGWSyw6gfGdR5o="}' http://localhost:8080/api/auth/revoke


---

- The SQLite database is persisted in the `data/` directory on your host.