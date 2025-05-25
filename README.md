
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"email":"konthalapalli@gmail.com", "password" : "abhishek"}'
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"email":"ajith@gmail.com", "password" : "abhishek"}'
curl -X POST http://localhost:8080/api/auth/signup -H "Content-Type: application/json" -d '{"email":"pravallika@gmail.com", "password" : "pravallika"}'

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtvbnRoYWxhcGFsbGlAZ21haWwuY29tIiwiZXhwIjoxNzQ4MjQ2MTU5LCJ1c2VyX2lkIjoxfQ.VaF38PebxUhMkgm9DCJskB9P5bmPcYa6zmU-6ThdIfc




curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImtvbnRoYWxhcGFsbGlAZ21haWwuY29tIiwiZXhwIjoxNzQ4MjQ2MTU5LCJ1c2VyX2lkIjoxfQ.VaF38PebxUhMkgm9DCJskB9P5bmPcYa6zmU-6ThdIfc" http://localhost:8080/api/protected

REFRESH REQUEST

cZYlveu-6G5t6bqwSnwKsGvsn2Q4BsGWSyw6gfGdR5o=


curl -X POST -H "Content-Type: application/json" -d '{"refresh_token" : "qYwUAHXdrbAGey7_hds6T8qYkmuI-vjnLiG9H7ODNWc="}' http://localhost:8080/api/auth/refresh




curl -X POST -H "Content-Type: application/json" -d '{"refresh_token" : "cZYlveu-6G5t6bqwSnwKsGvsn2Q4BsGWSyw6gfGdR5o="}' http://localhost:8080/api/auth/revoke