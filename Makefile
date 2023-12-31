path :=$(if $(path), $(path), "./")

up-rebuild: 
	@ docker-compose up --build -d

up: 
	@ docker-compose up -d

build-common:
	@ echo "selecting module $(app)"
	@ cd $(app) && go clean  
	@ cd $(app) && go mod tidy && go mod download
	@ cd $(app) && go mod verify

build: build-common
	@ echo clean
	@ rm -f $(app)/.bin/debug
	@ echo building...
	@ cd $(app) && go build -tags debug -o ".bin/debug" main.go
	@ ls -lah $(app)/.bin/debug

build-release: build-common
	@ echo clean
	@ rm -f $(app)/.bin/release
	@ echo build release
	@ cd $(app) && CGO_ENABLED=0 go build -ldflags='-w -s -extldflags "-static"' -a -o ".bin/release" main.go
	@ ls -lah $(app)/.bin/release


test: build-common
	@ cd $(app) && go test -v -cover ./...

scan:
	@ go install github.com/securego/gosec/v2/cmd/gosec@latest
	@ gosec -fmt=sarif -out=$(app).sarif -exclude=_test -severity=medium ./$(app)/... | 2>&1
	@ cat $(path)$(app).sarif


