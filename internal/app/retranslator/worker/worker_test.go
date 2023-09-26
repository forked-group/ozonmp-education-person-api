package worker

import (
	"context"
	"errors"
	mock_worker "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/worker/mocks"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func makeSuccessfulJobsChan(ctrl *gomock.Controller, n, times int) <-chan Job {
	job := mock_worker.NewMockJob(ctrl)
	job.EXPECT().Do().Times(times).Return(nil)

	ch := make(chan Job, n)

	for i := 0; i < n; i++ {
		ch <- job
	}

	return ch
}

func makeFailedJobsChan(ctrl *gomock.Controller, n, times int) <-chan Job {
	job := mock_worker.NewMockJob(ctrl)
	job.EXPECT().Do().Times(times).Return(errors.New("unknown error"))

	ch := make(chan Job, n)

	for i := 0; i < n; i++ {
		ch <- job
	}

	return ch
}

func makeCosedJobsChan(ctrl *gomock.Controller, _, times int) <-chan Job {
	job := mock_worker.NewMockJob(ctrl)
	job.EXPECT().Do().Times(times)

	ch := make(chan Job)
	close(ch)
	return ch
}

func TestConfig_Run(t *testing.T) {
	//lo.DebugEnable = true

	type fields struct {
		chLen    int
		jobTimes int
		In       func(ctrl *gomock.Controller, n, times int) <-chan Job
		Timeout  time.Duration
	}

	type args struct {
		ctxTimeout time.Duration
		name       string
	}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"[1] successful job three times (blocking on receive)",
			fields{
				3,
				3,
				makeSuccessfulJobsChan,
				2 * time.Second,
			},
			args{
				100 * time.Millisecond,
				"worker1",
			},
		},
		{
			"[2] failed jobs three times (blocking on receive)",
			fields{
				3,
				3,
				makeFailedJobsChan,
				10 * time.Millisecond,
			},
			args{
				100 * time.Millisecond,
				"worker2",
			},
		},
		{
			"[3] failed jobs three times (blocking on sleep)",
			fields{
				3,
				1,
				makeFailedJobsChan,
				2 * time.Second,
			},
			args{
				100 * time.Millisecond,
				"worker3",
			},
		},
		{
			"[4] closed input channel",
			fields{
				3,
				0,
				makeCosedJobsChan,
				2 * time.Second,
			},
			args{
				100 * time.Millisecond,
				"worker4",
			},
		},
		{
			"[5] context canceled",
			fields{
				3,
				0,
				makeSuccessfulJobsChan,
				2 * time.Second,
			},
			args{
				0,
				"worker1",
			},
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			cfg := &Config{
				In:      tt.fields.In(ctrl, tt.fields.chLen, tt.fields.jobTimes),
				Timeout: tt.fields.Timeout,
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

			tm := time.NewTimer(1 * time.Second)
			select {
			case <-tm.C:
				t.Error("blocking detected")
			case <-done:
			}
		})
	}
}
