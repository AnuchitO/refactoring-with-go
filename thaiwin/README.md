# Thaiwin API

3 route paths

~~~
/recently
    get currently checked in
/checkin
    check-in to any places
        input {
            "id": 1234,
            "place_id": 4321
        }
        output {
            "density": "ok"
        }
/checkout
        input {
            "id": 1234,
            "place_id": "4321"
        }
        output {}
~~~

3 tables in db

```go
type TablePeople struct {
	ID       int64
	MobileMo string
}

type TablePlaces struct {
	ID       int64
	Name     string
	Location Location
	Limit    int
}

type CacheVisitors struct {
	ID      int64
	PlaceID int64
}
```

## Technical Requirement

1. First version works
2. Second version is readable
3. Third version is testable
    - di by function
    - di by interface{}

- logging using uber-zap
- configuration using viper
- routing using gorilla-mux
- graceful shutting down
- distributed tracing using telemetry


## logging
`x-forwarded-for` in header please log it.

## note 
- mux.MiddlewareFunc กับ  http.Handler มัน compatability?

## How to run 
1. `make init` only once
2. `make run`