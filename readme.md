# KORESERVE #

#### How to create maria db Database ####
```
create database koreserve;
create database koreserve_iam;
create database koreserve_email;

grant all privileges on koreserve.* to 'gorm'@'localhost';
grant all privileges on koreserve_iam.* to 'gorm'@'localhost';
grant all privileges on koreserve_email.* to 'gorm'@'localhost';
```

#### How to create Swagger OpenApi Definition ####
swag cli설치
```
go install github.com/swaggo/swag/cmd/swag@latest
```

feature 디렉토리 안에서 실행
```
swag init -g openapi.go -output ../../../docs --parseDependency -d ./
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

