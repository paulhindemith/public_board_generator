// Code generated .* DO NOT EDIT.

package local

import (
	"testing"
	"context"
	"time"
	"golang.org/x/sync/errgroup"
	"github.com/paulhindemith/pkg-strategist/adapter"
)

func TestAdapter_Start(t *testing.T) {
	tests := []struct {
		name string
		f func(t *testing.T)
	}{
		{
			name: "bdd",
			f: func(t *testing.T){
				octx := adapter.TestContext("../../config/config.yaml", adapter.Jaeger, adapter.Prometheus)
				ctx, cancel := context.WithTimeout(octx, 1*time.Second)
				defer cancel()
				eg := errgroup.Group{}
				eg.Go(func()error{
					err := adapter.Main(ctx, NewAdapter)
					if err != nil {
						t.Fatal(err)
					}
					return err
				})
				<- ctx.Done()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.f)
	}
}
