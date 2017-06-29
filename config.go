package main

import (
	"flag"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

var cfg = struct {
	fs *flag.FlagSet `yaml:"-"`

	configFile string `yaml:"-"`

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

func init() {
	flag.StringVar(
		&cfg.configFile, "config.file", "nginxlog-exporter.yml",
		"Nginxlog exporter configuration file name.",
	)
	flag.Parse()

	readConfigFile()
}

func readConfigFile() {
	yamlFile, err := ioutil.ReadFile(cfg.configFile)
	if err != nil {
		log.Printf("failed to open %s : %s\n", cfg.configFile, err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Printf("failed to unmarshal %s : %s\n", cfg.configFile, err.Error())
	}
}
