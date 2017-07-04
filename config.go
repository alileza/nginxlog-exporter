package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var cfg = struct {
	fs *flag.FlagSet `yaml:"-"`

	configFile string `yaml:"-"`
	debug      bool   `yaml:"-"`

	Namespace     string `yaml:"namespace"`
	ListenAddress string `yaml:"listen_address"`
	MetricPath    string `yaml:"metric_path"`
	LogPath       string `yaml:"log_path"`
	LogFormat     string `yaml:"log_format"`

	Routes []string `yaml:"routes"`
}{
	Namespace:     "nginxlog",
	ListenAddress: ":9393",
	MetricPath:    "/metrics",
	LogPath:       "/var/log/nginx/access.log",
	LogFormat:     `$remote_addr - $remote_user [$time_local] "$method $endpoint $http_version" $status $body_bytes_sent "$http_referer" "$http_user_agent"`,
}

func debug(t string, d ...interface{}) {
	if !cfg.debug {
		return
	}
	fmt.Printf(t+"\n", d...)
}

func init() {
	flag.StringVar(
		&cfg.configFile, "config.file", "nginxlog-exporter.yml",
		"Nginxlog exporter configuration file name.",
	)
	flag.BoolVar(
		&cfg.debug, "debug", false,
		"Nginxlog exporter debug log.",
	)
	flag.Parse()

	readConfigFile()
}

func readConfigFile() {
	yamlFile, err := ioutil.ReadFile(cfg.configFile)
	if err != nil {
		debug("failed to open %s : %s\n", cfg.configFile, err.Error())
		return
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		debug("failed to unmarshal %s : %s\n", cfg.configFile, err.Error())
		return
	}

	debug("successfully load config : %+v", cfg)
}
