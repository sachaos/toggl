package util

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

type Writer interface {
	Write([]string) error
	Flush()
}

type TabWriter struct {
	w *tabwriter.Writer
}

func NewTabWriter(w io.Writer) *TabWriter {
	return &TabWriter{
		w: tabwriter.NewWriter(w, 0, 4, 1, ' ', 0),
	}
}

func (w *TabWriter) Flush() {
	w.w.Flush()
}

func (w *TabWriter) Write(record []string) error {
	string := strings.Join(record[:], "\t")
	fmt.Fprintln(w.w, string)
	return nil
}
