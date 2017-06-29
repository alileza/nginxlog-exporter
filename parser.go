package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/satyrius/gonx"
)

type parser struct {
	input         chan string
	parser        *gonx.Parser
	listenErrChan chan error
}

func NewParser(inp chan string) *parser {
	return &parser{
		input:  inp,
		parser: gonx.NewParser(cfg.LogFormat),
	}
}

func (p *parser) Run() {
	for {
		text := <-p.input
		parsedText, err := p.parser.ParseString(text)
		if err != nil {
			p.listenErrChan <- err
		}
		endpoint, _ := parsedText.Field("endpoint")
		endpoint = strings.Split(endpoint, "?")[0]
		code, _ := parsedText.Field("status")
		method, _ := parsedText.Field("method")
		bodyBytesSent, _ := parsedText.Field("body_bytes_sent")

		for _, route := range cfg.Routes {
			containRoutes := regexp.MustCompile(route)
			if containRoutes.MatchString(endpoint) {
				endpoint = route
			}
		}
		code = string(code[0]) + "xx"
		httpRequestCount.WithLabelValues(code, method, endpoint).Inc()

		bodyBytesSentF, _ := strconv.ParseFloat(bodyBytesSent, 64)
		httpResponseBodySizeBytes.WithLabelValues(code, method, endpoint).Add(bodyBytesSentF)

	}
}

func (p *parser) ListenError() <-chan error {
	return p.listenErrChan
}
