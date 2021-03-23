package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/olivere/elastic/v7"
)

//EsConn stores ES client
var EsConn *elastic.Client

//ConfigEs to configure ES client
type ConfigEs struct {
	URL string `env:"ES_URL"`
}

//GetEsConn returns the es client connexion.
func GetEsConn(ctx context.Context, foreverLoopDelay time.Duration) (*elastic.Client, error) {
	if EsConn == nil {
		if err := ConnectES(ctx, foreverLoopDelay); err != nil {
			return nil, fmt.Errorf("could not connect elastic search: %v", err)
		}
	}

	return EsConn, nil
}

//ConnectES creates a new ES client and stores it
func ConnectES(ctx context.Context, foreverLoopDelay time.Duration) error {
	cfg := ConfigEs{}
	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("could not parse env : %v", err)
	}

	client, err := elastic.NewClient(
		elastic.SetHealthcheck(true),
		elastic.SetSniff(false),
		elastic.SetURL(cfg.URL),
		elastic.SetHealthcheckInterval(15*time.Second),
	)
	if err != nil {
		return fmt.Errorf("could not create an elastic search client : %v", err)
	}

	for {
		delay := time.After(foreverLoopDelay)
		info, code, err := client.Ping(cfg.URL).Do(ctx)
		if err != nil {
			log.Printf("%v... retrying", err)
		} else {
			log.Printf("Elasticsearch returned with code %d and version %s\n", code,
				info.Version.Number)
			EsConn = client
			log.Printf("connected to es client")

			break
		}
		<-delay
	}

	return nil

}
