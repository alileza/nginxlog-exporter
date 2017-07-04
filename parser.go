package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/satyrius/gonx"
)

type Parser struct {
	input         chan string
	parser        *gonx.Parser
	listenErrChan chan error
}

func NewParser(inp chan string) *Parser {
	return &Parser{
		input:         inp,
		parser:        gonx.NewParser(cfg.LogFormat),
		listenErrChan: make(chan error),
	}
}

func (p *Parser) Run() {
	for {
		text := <-p.input
		parsedText, err := p.parser.ParseString(text)
		if err != nil {
			p.listenErrChan <- err
			return
		}
		endpoint, _ := parsedText.Field("endpoint")
		endpoint = strings.Split(endpoint, "?")[0]
		code, _ := parsedText.Field("status")
		method, _ := parsedText.Field("method")
		bodyBytesSent, _ := parsedText.Field("body_bytes_sent")

		debug("parsed endpoint=%s code=%s method=%s bodyBytesSent=%s", endpoint, code, method, bodyBytesSent)
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

func (p *Parser) ListenError() <-chan error {
	return p.listenErrChan
}
