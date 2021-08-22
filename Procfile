web: go build -o bin/upcourse -v .
web: bin/upcourse
release: go build -o bin/migrate -v database/migrate
release: bin/migrate
release: go build -o bin/seed -v database/seed
release: bin/seed