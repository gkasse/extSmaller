OS := darwin
ARCHS := 386 amd64
depExist := `which dep`

# extSmaller use to dep to resolve libraries, therefore installed dep.
init:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

test:
	go test -cover `go list ./...`

ci-test:
	install -D /dev/null ./out/coverage.txt; \
	for d in `go list ./...`; do \
			go test -coverprofile=profile.out -v $$d; \
			if [ -f profile.out ]; then \
					cat profile.out >> ./out/coverage.txt; \
					rm profile.out; \
			fi; \
	done ; \
	go tool cover -html=./out/coverage.txt -o=./out/coverage.html

build:
	for arch in $(ARCHS) ; do \
		GOOS=darwin GOARCH=$$arch go build -o ./out/extSmaller-darwin-$$arch extSmaller.go ; \
	done ; \
	for arch in $(ARCHS) ; do \
		GOOS=windows GOARCH=$$arch go build -o ./out/extSmaller-windows-$$arch.exe extSmaller.go ; \
	done
