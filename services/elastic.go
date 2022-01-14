package services

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
	"github.com/projectdiscovery/gologger"
	"github.com/wahyuhadi/ESgo/es"
	"github.com/wahyuhadi/semgrep-integrator/models"
)

// Model data

func Elastic(opts *models.Options) {

	cfg := elasticsearch.Config{
		Addresses: []string{opts.ElasticUrl},
		Username:  opts.ElasticUser, // if ES need this
		Password:  opts.ElasticPass,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
				// ...
			},
		},
	}
	c, _ := elasticsearch.NewClient(cfg)
	// PushexamplePushData(c)
	pushdata(opts, c)

}

func pushdata(opts *models.Options, c *elasticsearch.Client) {
	var semgrep models.Semgrep
	err := json.NewDecoder(os.Stdin).Decode(&semgrep)
	if err != nil {
		gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error json decoder")
		return
	}

	// parsing with esutil from elastic
	for _, datas := range semgrep.Results {
		data := esutil.NewJSONReader(&datas)
		// Push data to elastic
		response, err := es.PushData(c, opts.ElasicIndex, data)
		if err != nil {
			gologger.Info().Str("Error", fmt.Sprintf("%v", err.Error())).Msg("Error push data")

		}
		gologger.Info().Str("Is Error ", fmt.Sprintf("%s", response.IsError())).Msg("Success Push data to elastic")

	}

}
