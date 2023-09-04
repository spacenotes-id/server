all: start

build:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/spacenotes-server .

install:
	go install .

start: db-start build
	./bin/spacenotes-server $(ARGS)

dev: db-start
	go run . $(ARGS)

# dependency:
# - watchexec = https://github.com/watchexec/watchexec
dev-watch: db-start
	watchexec -c -r -e go -- go run . $(ARGS)

sqlc-generate:
	sqlc generate

swag-fmt:
	swag fmt

swag-init:
	swag init

swag: swag-fmt swag-init

# test:
# 	go test -v ./...
#
# test-cover:
# 	go test -v -cover ./...
#
# test-cover-watch:
# 	watchexec -c -r -e go -- go test -v -cover ./...
#
# test-cover-html:
# 	go test -coverprofile cover.out ./... && \
# 		go tool cover -html=cover.out
#
# generate:
# 	go generate ./...

clean:
	rm -f ./bin/* ./cover.out

db-start:
	systemctl is-active postgresql || systemctl start postgresql

db-stop:
	systemctl is-active postgresql && systemctl stop postgresql

db-status:
	systemctl status postgresql

.PHONY: build install start dev dev-watch sqlc-generate test test-cover test-cover-watch test-cover-html generate clean db-start db-stop db-status, swag-init, swag-fmt, swag
