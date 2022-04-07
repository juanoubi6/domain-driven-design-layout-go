package sql

import (
	"domain-driven-design-layout/infrastructure/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"sync"
)

var registerQueryHistogram sync.Once

// QueryTimeHistogram registers query time to measure db access times
var QueryTimeHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "query_time",
	Help:    "Query execution time in seconds",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2, 3, 5, 10},
}, []string{"query_name"})

const connStringTemplate = "host=%s user=%s dbname=%s password=%s port=%d sslmode=disable"

func CreateDatabaseConnection(config config.SQLConfig) *sqlx.DB {
	uri := fmt.Sprintf(
		connStringTemplate,
		config.Host,
		config.UserName,
		config.DbName,
		config.Password,
		config.Port,
	)

	db, err := sqlx.Connect("postgres", uri)
	if err != nil {
		log.Fatalln("DB | Could not connect to the database: " + err.Error())
	}

	registerQueryHistogram.Do(func() {
		prometheus.MustRegister(QueryTimeHistogram)
	})

	return db
}
