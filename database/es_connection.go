package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/olivere/elastic/v7"
)

//EsConn stores ES client
var EsConn *elastic.Client

//ConfigEs to configure ES client
type ConfigEs struct {
	URL string
}

//ConnectES creates a new ES client and stores it
func ConnectES(ctx context.Context, cfg *ConfigEs, foreverLoopDelay time.Duration) error {

	client, err := elastic.NewClient(
		elastic.SetHealthcheck(true),
		elastic.SetSniff(true),
		elastic.SetURL(cfg.URL),
		elastic.SetHealthcheckInterval(15*time.Second),
	)
	if err != nil {
		return fmt.Errorf("could not create an eventsourcing client : %v", err)
	}

	for {
		delay := time.After(foreverLoopDelay)
		info, code, err := client.Ping(cfg.URL).Do(ctx)
		if err != nil {
			log.Printf("%v... retrying", err)
		} else {
			fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
			EsConn = client
			log.Printf("connected to es client")

			break
		}
		<-delay
	}

	return nil

}
