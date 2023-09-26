package producer

import (
	"context"
	"errors"
	mock_producer "github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator/producer/mocks"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestConfig_Run(t *testing.T) {
	//lo.DebugEnable = true

	makeChanSize := func(n int) chan *person.PersonEvent {
		ch := make(chan *person.PersonEvent, n)
		for i := 0; i < n; i++ {
			ch <- &person.PersonEvent{ID: uint64(i + 1)}
		}
		return ch
	}

	makeCloseChan := func() chan *person.PersonEvent {
		ch := make(chan *person.PersonEvent)
		close(ch)
		return ch
	}

	type fields struct {
		Timeout time.Duration
		In      <-chan *person.PersonEvent
		Sender  func(ctrl *gomock.Controller) EventSender
		Ok      chan<- uint64
		Error   chan<- uint64
	}

	type args struct {
		ctxTimeout time.Duration
		name       string
	}

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantOkLen    int
		wantErrorLen int
	}{
		{
			"[1] sending three times (blocking on receiving)",
			fields{
				2 * time.Second,
				makeChanSize(3),
				func(ctrl *gomock.Controller) EventSender {
					sender := mock_producer.NewMockEventSender(ctrl)
					sender.EXPECT().Send(gomock.Any()).Times(3).Return(nil)
					return sender
				},
				make(chan uint64, 10),
				make(chan uint64, 10),
			},
			args{
				100 * time.Millisecond,
				"producer1",
			},
			3,
			0,
		},
		{
			"[2] sending error (blocking on sleep)",
			fields{
				2 * time.Second,
				makeChanSize(3),
				func(ctrl *gomock.Controller) EventSender {
					sender := mock_producer.NewMockEventSender(ctrl)
					sender.EXPECT().Send(gomock.Any()).Times(1).
						Return(errors.New("unknown error"))
					return sender
				},
				make(chan uint64, 10),
				make(chan uint64, 10),
			},
			args{
				100 * time.Millisecond,
				"producer2",
			},
			0,
			1,
		},
		{
			"[3] sending error three times (blocking on send last report)",
			fields{
				10 * time.Millisecond,
				makeChanSize(3),
				func(ctrl *gomock.Controller) EventSender {
					sender := mock_producer.NewMockEventSender(ctrl)
					sender.EXPECT().Send(gomock.Any()).Times(3).
						Return(errors.New("unknown error"))
					return sender
				},
				make(chan uint64, 10),
				make(chan uint64, 2),
			},
			args{
				100 * time.Millisecond,
				"producer3",
			},
			0,
			2,
		},
		{
			"[4] closed input channel",
			fields{
				2 * time.Second,
				makeCloseChan(),
				func(ctrl *gomock.Controller) EventSender {
					sender := mock_producer.NewMockEventSender(ctrl)
					sender.EXPECT().Send(gomock.Any()).Times(0)
					return sender
				},
				make(chan uint64, 10),
				make(chan uint64, 10),
			},
			args{
				100 * time.Millisecond,
				"producer4",
			},
			0,
			0,
		},
		{
			"[5] canceled context",
			fields{
				2 * time.Second,
				makeCloseChan(),
				func(ctrl *gomock.Controller) EventSender {
					sender := mock_producer.NewMockEventSender(ctrl)
					sender.EXPECT().Send(gomock.Any()).Times(0)
					return sender
				},
				make(chan uint64, 10),
				make(chan uint64, 10),
			},
			args{
				0,
				"producer5",
			},
			0,
			0,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			cfg := &Config{
				Timeout: tt.fields.Timeout,
				In:      tt.fields.In,
				Sender:  tt.fields.Sender(ctrl),
				Ok:      tt.fields.Ok,
				Error:   tt.fields.Error,
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

			if got := len(tt.fields.Ok); got != tt.wantOkLen {
				t.Errorf("got len(Ok)=%d, want %d", got, tt.wantOkLen)
			}

			if got := len(tt.fields.Error); got != tt.wantErrorLen {
				t.Errorf("got len(Error)=%d, want %d", got, tt.wantErrorLen)
			}
		})
	}
}
