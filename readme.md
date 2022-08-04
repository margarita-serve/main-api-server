# KORESERVE #

#### How to create maria db Database ####
```
create database koreserve;

grant all privileges on koreserve.* to 'gorm'@'localhost';

```

#### How to create Swagger OpenApi Definition ####
swag cli설치
```
go install github.com/swaggo/swag/cmd/swag@latest
```

feature 디렉토리 안에서 실행
```
$GOPATH/bin/swag init -g ./interface/restapi/feature/openapi.go -output ./docs --parseDependency -d ./
```
```
http://localhost:32022/openapi/swagger/index.html
```

#### How to Run  ####
```
go run main.go server restapi
```
connect URL is
http://localhost:32022

#### How to Migrate DB  ####
```
go run main.go db migrate
```

