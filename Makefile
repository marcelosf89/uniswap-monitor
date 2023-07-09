path :=$(if $(path), $(path), "./")

up-rebuild: 
	@ docker-compose up --build

up: 
	@ docker-compose up

build-common:
	@ echo "selecting module $(app)"
	@ cd $(app) && go clean  
	@ cd $(app) && go mod tidy && go mod download
	@ cd $(app) && go mod verify

build: build-common
	@ echo clean
	@ rm -f $(path).bin/debug/$(app)
	@ echo building...
	@ cd $(app) && go build -tags debug -o "$(path).bin/debug/$(app)" main.go
	@ ls -lah $(path).bin/debug/$(app)

build-release: build-common
	@ echo clean
	@ rm -f $(path).bin/release/$(app)
	@ echo build release
	@ cd $(app) && CGO_ENABLED=0 go build -ldflags='-w -s -extldflags "-static"' -a -o "$(path).bin/release/$(app)" main.go
	@ ls -lah $(path).bin/release/$(app)


test: build-common
	@ cd $(path)/$(app) 
	@ go test -v -cover $(app)/...
	@ cd ..


scan:
	@ go install github.com/securego/gosec/v2/cmd/gosec@latest
	@ gosec -fmt=sarif -out=$(app).sarif -exclude=_test -severity=medium ./$(app)/... 
	@ echo ""
	@ cat $(path)$(app).sarif


