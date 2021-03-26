init:
	sqlite3 thaiwin.db < _scripts/init.sql

run:
	PORT=8000 DB_CONN=thaiwin.db go run main.go

test: 
	go test -coverprofile thaiwin.html ./...
	go tool  cover -html=thaiwin.html
