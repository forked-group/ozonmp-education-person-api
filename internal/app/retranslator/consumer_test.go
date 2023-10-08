package retranslator

import (
	"context"
	"errors"
	mock_retranslator "github.com/aaa2ppp/ozonmp-education-person-api/internal/app/retranslator/mocks"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestConfig_Run(t *testing.T) {
	type fields struct {
		BatchSize int
		Timeout   time.Duration
		Repo      func(ctrl *gomock.Controller) EventLocker
		Out       chan<- []event
	}

	type args struct {
		ctxTimeout time.Duration
	}

	batch := []event{{ID: 1}, {ID: 2}, {ID: 3}}

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"[1] locking three times (blocking on last send)",
			fields{
				3,
				10 * time.Second,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_retranslator.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(3))).
						Times(3).
						Return(batch, nil)
					return repo
				},
				make(chan []event, 2),
			},
			args{
				100 * time.Millisecond,
			},
		},
		{
			"[2] error locking (blocking on sleep)",
			fields{
				10,
				2 * time.Second,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_retranslator.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(10))).
						Times(1).
						Return(batch, errors.New("unknown error"))
					return repo
				},
				make(chan []event, 10),
			},
			args{
				100 * time.Millisecond,
			},
		},
		{
			"[3] incomplete batch (blocking on sleep)",
			fields{
				10,
				2 * time.Second,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_retranslator.NewMockEventLocker(ctrl)
					repo.EXPECT().
						Lock(gomock.Eq(uint64(10))).
						Return(batch, nil).
						Times(1)
					return repo
				},
				make(chan []event, 1),
			},
			args{
				100 * time.Millisecond,
			},
		},
		{
			"[4] incomplete batch two times (blocking on last send)",
			fields{
				10,
				10 * time.Millisecond,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_retranslator.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(10))).
						Times(2).
						Return(batch, nil)
					return repo
				},
				make(chan []event, 1),
			},
			args{
				100 * time.Millisecond,
			},
		},
		{
			"[5] canceled context",
			fields{
				10,
				2 * time.Millisecond,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_retranslator.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(10))).
						Times(0)
					return repo
				},
				make(chan []event, 1),
			},
			args{
				0,
			},
		},
		// TODO: AddField test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			cfg := &consumerConfig{
				BatchSize: tt.fields.BatchSize,
				Timeout:   tt.fields.Timeout,
				Repo:      tt.fields.Repo(ctrl),
				Out:       tt.fields.Out, // TODO: Fix, now: Out chan<- []event. Fix what???
			}

			var ctx context.Context
			var cancel, sendTerm context.CancelFunc

			if tt.args.ctxTimeout > 0 {
				ctx, cancel = context.WithTimeout(context.Background(),
					tt.args.ctxTimeout)
				ctx, sendTerm = contextWithTerm(ctx)
				defer func() {
					sendTerm()
					cancel()
				}()
			} else {
				ctx, cancel = context.WithCancel(context.Background())
				ctx, sendTerm = contextWithTerm(ctx)
				sendTerm()
				cancel()
			}

			done := make(chan struct{})
			go func() {
				defer close(done)
				cfg.Run(ctx)
			}()

			tm := time.NewTimer(1 * time.Second)
			select {
			case <-done:
			case <-tm.C:
				t.Error("blocking detected")
			}
		})
	}
}
