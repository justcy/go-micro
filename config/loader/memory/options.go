package memory

import (
	"go-micro.kanter.cn/v1/config/loader"
	"go-micro.kanter.cn/v1/config/reader"
	"go-micro.kanter.cn/v1/config/source"
)

// WithSource appends a source to list of sources.
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader.
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}

func WithWatcherDisabled() loader.Option {
	return func(o *loader.Options) {
		o.WithWatcherDisabled = true
	}
}
