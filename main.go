package main

import (
	"github.com/jenningsloy318/hana_exporter/collector"
	"github.com/jenningsloy318/hana_exporter/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
	"os"
)

// define  flag
var (
	listenAddress = kingpin.Flag(
		"web.listen-address",
		"Address to listen on for web interface and telemetry.",
	).Default(":9460").String()
	logLevel = kingpin.Flag("log.level", "Set log level").Default("info").String()
	c        config.Config
)

// scrapers lists all possible collection methods and if they should be enabled by default.
var scrapers = map[collector.Scraper]bool{
	collector.ScrapeHostResourceUtilization{}: true,
	collector.ScrapeServiceStatistics{}:       true,
	collector.ScrapeLicenseStatus{}:           true,
	collector.ScrapeDisks{}:                   true,
	collector.ScrapeSharedMemory{}:            true,
	collector.ScrapeCsTables{}:                true,
	collector.ScrapeServiceReplication{}:      true,
	collector.ScrapeSystemConfig{}:            true,
	collector.ScrapeSystemReplication{}:       true,
	collector.ScrapeCsUnloads{}:               true,
	collector.ScrapeCsLoads{}:                 true,
	collector.ScrapeRsTables{}:                true,
}

// define new http handleer
func newHandler(scrapers []collector.Scraper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		registry := prometheus.NewRegistry()
		registry.MustRegister(collector.New(c, scrapers))

		// Remove Go collector
		prometheus.Unregister(prometheus.NewGoCollector())

		gatherers := prometheus.Gatherers{
			prometheus.DefaultGatherer,
			registry,
		}
		// Delegate http serving to Prometheus client library, which will call collector.Collect.
		h := promhttp.HandlerFor(gatherers, promhttp.HandlerOpts{})
		h.ServeHTTP(w, r)
	}
}

func main() {
	// Generate ON/OFF flags for all scrapers.
	scraperFlags := map[collector.Scraper]*bool{}
	for scraper, enabledByDefault := range scrapers {
		defaultOn := "false"
		if enabledByDefault {
			defaultOn = "true"
		}

		f := kingpin.Flag(
			"collect."+scraper.Name(),
			scraper.Help(),
		).Default(defaultOn).Bool()

		scraperFlags[scraper] = f
	}

	c = config.Config{
		Databases: config.DatabaseConfig{
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
			User:     os.Getenv("USER"),
			Password: os.Getenv("PASS"),
			Timeout:  os.Getenv("TIMEOUT"),
		},
	}

	// Parse flags.
	kingpin.Version(version.Print("hana_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.SetLevel(parseLogLevel(*logLevel))
	log.Debugf("database config: %+v", c.Databases)

	// landingPage contains the HTML served at '/'.
	// TODO: Make this nicer and more informative.
	var landingPage = []byte(`<html>
<head><title>HANA exporter</title></head>
<body>
<h1>HANA exporter</h1>
<p><a href='/metrics'>Metrics</a></p>
</body>
</html>
`)

	log.Infoln("Starting hana_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())

	// Register only scrapers enabled by flag.
	log.Infof("Enabled scrapers:")
	var enabledScrapers []collector.Scraper
	for scraper, enabled := range scraperFlags {
		if *enabled {
			log.Infof(" --collect.%s", scraper.Name())
			enabledScrapers = append(enabledScrapers, scraper)
		}
	}

	http.HandleFunc("/metrics", newHandler(enabledScrapers))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(landingPage)
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func parseLogLevel(value string) log.Level {
	level, err := log.ParseLevel(value)

	if err != nil {
		log.WithField("log-level-value", value).Warningln("invalid log level from env var, using info")
		return log.ErrorLevel
	}

	return level
}
