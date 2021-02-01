package database

import (
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
func ConnectES(cfg *ConfigEs) error {

	client, err := elastic.NewClient(
		elastic.SetSniff(true),
		elastic.SetURL(cfg.URL),
		elastic.SetHealthcheckInterval(15*time.Second),
	)
	if err != nil {
		return fmt.Errorf("could not create an es client : %v", err)
	}

	log.Printf("connected to es client")

	EsConn = client

	return nil
}
