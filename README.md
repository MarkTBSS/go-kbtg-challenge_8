```
go get -u github.com/swaggo/swag
docker-compose up
go run .

http://localhost:1323/api/v1/wallets
http://localhost:1323/swagger/index.html

cd wallet
go test -v -cover 

docker build --no-cache -t aiyaraaiya/go-kbtg-challenge_8:v1.0 .
docker run -p 1324:1323 aiyaraaiya/go-kbtg-challenge_8:v1.0
```