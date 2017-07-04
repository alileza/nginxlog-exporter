package main

import "github.com/hpcloud/tail"

type Tailer struct {
	listenErrChan chan error
	Out           chan string
}

func NewTailer() *Tailer {
	return &Tailer{
		listenErrChan: make(chan error),
		Out:           make(chan string),
	}
}

func (t *Tailer) Run() {
	file, err := tail.TailFile(cfg.LogPath, tail.Config{Follow: true})
	if err != nil {
		t.listenErrChan <- err
		return
	}

	for line := range file.Lines {
		t.Out <- line.Text
	}
}

func (t *Tailer) ListenError() <-chan error {
	return t.listenErrChan
}
