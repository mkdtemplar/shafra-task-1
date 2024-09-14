To run the application in terminal from the root of the application:
go run cmd/main.go

For migration of the database migrate package needs to be install.
For instructions on how to install the migration package go to: https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

This application uses Gin framework for HTTP requests, to install the Gin framework run in terminal:
go get -u github.com/gin-gonic/gin

For database manipulation i.e. GORM framework is used to install use this command in terminal:
go get -u gorm.io/gorm

Application uses the viper package for handling the environment variables, to install the viper package type in terminal:
go get github.com/spf13/viper

To run test type in terminal: make tests