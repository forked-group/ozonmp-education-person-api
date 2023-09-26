package collector

import (
	"context"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/worker"
	"testing"
	"time"
)

type Job = worker.Job

func TestConfig_Run(t *testing.T) {
	//lo.DebugEnable = true

	makeChanSize := func(n int) <-chan uint64 {
		ch := make(chan uint64, n)
		for i := 1; i <= n; i++ {
			ch <- uint64(i)
		}
		return ch
	}

	makeClosedChan := func(n int) <-chan uint64 {
		ch := make(chan uint64, n)
		close(ch)
		return ch
	}

	type fields struct {
		Job          func([]uint64) error
		inLen        int
		In           func(n int) <-chan uint64
		Out          chan<- Job
		BatchSize    int
		FlushTimeout time.Duration
	}
	type args struct {
		ctxTimeout time.Duration
		name       string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutLen int
	}{
		{
			"[1] ten times",
			fields{
				nil,
				10,
				makeChanSize,
				make(chan Job, 10),
				3,
				20 * time.Millisecond,
			},
			args{
				100 * time.Millisecond,
				"collector1",
			},
			(10 + 3 - 1) / 3,
		},
		{
			"[2] closed input channel",
			fields{
				nil,
				10,
				makeClosedChan,
				make(chan Job, 10),
				3,
				20 * time.Millisecond,
			},
			args{
				100 * time.Millisecond,
				"collector2",
			},
			0,
		},
		{
			"[3] empty input channel",
			fields{
				nil,
				0,
				makeChanSize,
				make(chan Job, 10),
				3,
				20 * time.Millisecond,
			},
			args{
				100 * time.Millisecond,
				"collector3",
			},
			0,
		},
		{
			"[4] context canceled",
			fields{
				nil,
				10,
				makeChanSize,
				make(chan Job, 10),
				3,
				20 * time.Millisecond,
			},
			args{
				0,
				"collector4",
			},
			0,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &Config{
				Job:       tt.fields.Job,
				In:        tt.fields.In(tt.fields.inLen),
				Out:       tt.fields.Out,
				BatchSize: tt.fields.BatchSize,
				Timeout:   tt.fields.FlushTimeout,
			}

			var ctx context.Context
			var cancel context.CancelFunc

			if tt.args.ctxTimeout > 0 {
				ctx, cancel = context.WithTimeout(context.Background(),
					tt.args.ctxTimeout)
				defer cancel()
			} else {
				ctx, cancel = context.WithCancel(context.Background())
				cancel()
			}

			done := make(chan struct{})
			go func() {
				defer close(done)

				cfg.Run(ctx, tt.args.name)
			}()

			tm := time.NewTimer(200 * time.Millisecond)
			select {
			case <-tm.C:
				t.Error("blocking detected")
			case <-done:
			}

			if got := len(cfg.Out); got != tt.wantOutLen {
				t.Errorf("got len(Out)=%d, want %d", got, tt.wantOutLen)
			}
		})
	}
}
