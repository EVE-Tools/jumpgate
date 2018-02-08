package main

import (
	"flag"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/EVE-Tools/jumpgate/element43/services/esiMarkets"
	"github.com/EVE-Tools/jumpgate/element43/services/marketStats"
	"github.com/EVE-Tools/jumpgate/element43/services/staticData"
	"github.com/EVE-Tools/jumpgate/element43/services/topStations"
)

// Config holds the application's configuration info from the environment.
type Config struct {
	LogLevel        string `default:"info" envconfig:"log_level"`
	ListenHost      string `default:":8000" split_words:"true"`
	ESIMarketsHost  string `default:"esi-markets.element43.svc.cluster.local:43000" envconfig:"esi_markets_host"`
	MarketStatsHost string `default:"market-stats.element43.svc.cluster.local:43000" envconfig:"market_stats_host"`
	StaticDataHost  string `default:"static-data.element43.svc.cluster.local:43000" envconfig:"static_data_host"`
	TopStationsHost string `default:"top-stations.element43.svc.cluster.local:43000" envconfig:"top_stations_host"`
}

// Load configuration from environment
func loadConfig() Config {
	config := Config{}
	envconfig.MustProcess("JUMPGATE", &config)

	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}

	logrus.SetLevel(logLevel)
	logrus.Debugf("Config: %q", config)
	return config
}

func run() error {
	config := loadConfig()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(50000000)),
		grpc.WithInsecure(),
	}

	err := esiMarkets.RegisterESIMarketsHandlerFromEndpoint(ctx, mux, config.ESIMarketsHost, opts)
	if err != nil {
		return err
	}

	err = marketStats.RegisterMarketStatsHandlerFromEndpoint(ctx, mux, config.MarketStatsHost, opts)
	if err != nil {
		return err
	}

	err = staticData.RegisterStaticDataHandlerFromEndpoint(ctx, mux, config.StaticDataHost, opts)
	if err != nil {
		return err
	}

	err = topStations.RegisterTopStationsHandlerFromEndpoint(ctx, mux, config.TopStationsHost, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(config.ListenHost, mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
