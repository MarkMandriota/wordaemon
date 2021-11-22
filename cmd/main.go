// Copyright 2021 Mark Mandriota. All right reserved.
// Use of this source code is governed by MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/MarkMandriota/wordaemon/pkg/words"
	"golang.design/x/clipboard"
)

var logger = log.New(os.Stderr, "# ", log.Ltime)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dict := make(words.Dict)
	words.LoadDict(os.Stdin, dict)

	logger.Println("starting daemon ...")

	go clipboardHandler(ctx, dict)
	defer clipboard.Write(0, nil)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)

	<-sigc
}

func clipboardHandler(ctx context.Context, dict words.Dict) {
	rand.Seed(int64(time.Now().Nanosecond()) * int64(os.Getpid()))

	ticker := time.NewTicker(time.Second / 3)
	defer ticker.Stop()

	waiting := false

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			text := clipboard.Read(clipboard.FmtText)

			if r, size := utf8.DecodeRune(text); !waiting {
				switch r {
				case ']':
					waiting = true
				case '~':
					if t, err := time.ParseDuration(string(text[size:])); err == nil {
						ticker.Reset(t)
					}

					clipboard.Write(0, nil)
				case '-':
					syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
				default:
					clipboard.Write(clipboard.FmtText,
						dict.Choice(clipboard.Read(clipboard.FmtText)))
				}
			} else if r == '[' {
				waiting = false
			}
		}
	}
}
