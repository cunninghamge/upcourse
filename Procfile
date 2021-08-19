web: go build -o bin/upcourse -v .
web: bin/upcourse
release: go build -o bin/migrate -v database/migrate/main.go
release: bin/migrate
elease: go build -o bin/seed -v database/seed/main.go
release: bin/seed