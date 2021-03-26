package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func init() {
	viper.SetDefault("port", "8000")
	viper.SetDefault("db.conn", "thaiwin.db")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	hostname, _ := os.Hostname()
	logger = logger.With(zap.String("hostname", hostname))
	zap.ReplaceGlobals(logger)

	db, err := sql.Open("sqlite3", viper.GetString("db.conn"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()
	r := mux.NewRouter()
	r.Use(LoggerMiddleware(logger))

	r.HandleFunc("/recently", Recently).Methods(http.MethodPost)
	r.Handle("/checkin", &CheckIn{InsertCheckIn: insertCheckIn}).Methods(http.MethodPost)
	r.HandleFunc("/checkout", CheckOut).Methods(http.MethodPost)

	port := viper.GetString("port")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	zap.L().Info("starting...", zap.String("port", port))
	log.Fatal(srv.ListenAndServe())
}

type Check struct {
	ID      int64 `json:"id"`
	PlaceID int64 `json:"place_id"`
}

type Location struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// Recently returns currently visited
func Recently(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

type CheckIn struct {
	InsertCheckIn InsertCheckIn
}

type InsertCheckIn func(id, placeID int64) error

func insertCheckIn(id, placeID int64) error {
	db, err := sql.Open("sqlite3", "thaiwin.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO visits VALUES(?, ?);", id, placeID)
	return err
}

// CheckIn check-in to place, returns density (ok, too much)
func (c *CheckIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Context().Value("logger").(*zap.Logger).Info("check-in")
	chk := Check{}
	if err := json.NewDecoder(r.Body).Decode(&chk); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	defer r.Body.Close()

	err := c.InsertCheckIn(chk.ID, chk.PlaceID)
	if err != nil {
		r.Context().Value("logger").(*zap.Logger).Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "density": "ok" }`))
}


// CheckOut check-out from place
func CheckOut(w http.ResponseWriter, r *http.Request) {

}

func LoggerMiddleware(logger *zap.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l := logger.With(zap.String("middleware", "some data in middleware"))
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "logger", l)))
		})
	}
}
