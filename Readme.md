## Auth Service
This service provides authentication and authorization functionalities for applications.
It includes features such as user registration, login, token generation, 
and validation using JWT (JSON Web Tokens). 
The service is built using the Go programming language and leverages 
the Gin framework for handling HTTP requests,
GORM for database interactions, and Docker for containerization.

## Dependencies

```
go module github/com/Gajju8989/Auth_Service
go 1.23.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.10.0
	github.com/go-sql-driver/mysql v1.9.0
	github.com/google/uuid v1.6.0
	github.com/google/wire v0.6.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.34.0
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/bytedance/sonic v1.11.6 // indirect
	github.com/bytedance/sonic/loader v0.1.1 // indirect
	github.com/cloudwego/base64x v0.1.4 // indirect
	github.com/cloudwego/iasm v0.2.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.20.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.7 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.2.2 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	golang.org/x/arch v0.8.0 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

```

## To run the server
``` terminal
 docker-compose -f docker-compose.yml build
 docker-compose -f docker-compose.yml up -d
```


## API Documentation

### Postman API Documentation
Here's the link to the API documentation for HTTP endpoints:
- [Postman API Documentation](https://documenter.getpostman.com/view/29203481/2sAYdcsCT5#3710c810-dba4-4267-a0a2-f8fe788a34bc)
