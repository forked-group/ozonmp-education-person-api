package consumer

import (
	"context"
	"errors"
	mock_consumer "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/consumer/mocks"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestConfig_Run(t *testing.T) {
	//lo.DebugEnable = true

	type fields struct {
		BatchSize int
		Timeout   time.Duration
		Repo      func(ctrl *gomock.Controller) EventLocker
		Out       chan<- *person.PersonEvent
	}

	type args struct {
		ctxTimeout time.Duration
		name       string
	}

	batch := []person.PersonEvent{{ID: 1}, {ID: 2}, {ID: 3}}

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
					repo := mock_consumer.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(3))).
						Times(3).
						Return(batch, nil)
					return repo
				},
				make(chan *person.PersonEvent, len(batch)*3-1),
			},
			args{
				100 * time.Millisecond,
				"consumer1",
			},
		},
		{
			"[2] error locking (blocking on sleep)",
			fields{
				10,
				2 * time.Second,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_consumer.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(10))).
						Times(1).
						Return(batch, errors.New("unknown error"))
					return repo
				},
				make(chan *person.PersonEvent, 10),
			},
			args{
				100 * time.Millisecond,
				"consumer2",
			},
		},
		{
			"[3] incomplete batch (blocking on sleep)",
			fields{
				10,
				2 * time.Second,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_consumer.NewMockEventLocker(ctrl)
					repo.EXPECT().
						Lock(gomock.Eq(uint64(10))).
						Return(batch, nil).
						Times(1)
					return repo
				},
				make(chan *person.PersonEvent, 3),
			},
			args{
				100 * time.Millisecond,
				"consumer3",
			},
		},
		{
			"[4] incomplete batch two times (blocking on last send)",
			fields{
				10,
				10 * time.Millisecond,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_consumer.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(10))).
						Times(2).
						Return(batch, nil)
					return repo
				},
				make(chan *person.PersonEvent, len(batch)*2-1),
			},
			args{
				100 * time.Millisecond,
				"consumer4",
			},
		},
		{
			"[5] canceled context",
			fields{
				10,
				2 * time.Millisecond,
				func(ctrl *gomock.Controller) EventLocker {
					repo := mock_consumer.NewMockEventLocker(ctrl)
					repo.EXPECT().Lock(gomock.Eq(uint64(10))).
						Times(0)
					return repo
				},
				make(chan *person.PersonEvent, 3),
			},
			args{
				0,
				"consumer4",
			},
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			cfg := &Config{
				BatchSize: tt.fields.BatchSize,
				Timeout:   tt.fields.Timeout,
				Repo:      tt.fields.Repo(ctrl),
				Out:       tt.fields.Out, // TODO: Fix, now: Out chan<- []person.PersonEvent
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
			case <-done:
			case <-tm.C:
				t.Error("blocking detected")
			}
		})
	}
}
