set -xe

go get github.com/gin-gonic/gin

go get gopkg.in/go-playground/validator.v10

go build -o bin/application main.go

