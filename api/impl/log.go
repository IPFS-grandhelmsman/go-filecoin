package impl

import (
	"context"
	"io"

	writer "gx/ipfs/QmRREK2CAZ5Re2Bd9zZFG6FeYDppUWt5cMgsoUEp3ktgSr/go-log/writer"
)

type nodeLog struct {
	api *nodeAPI
}

func newNodeLog(api *nodeAPI) *nodeLog {
	return &nodeLog{api: api}
}

func (api *nodeLog) Tail(ctx context.Context) io.Reader {
	r, w := io.Pipe()
	go func() {
		defer w.Close() // nolint: errcheck
		<-ctx.Done()
	}()

	writer.WriterGroup.AddWriter(w)

	return r
}