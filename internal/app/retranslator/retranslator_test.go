package retranslator_test

import (
	"context"
	"errors"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/retranslator"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	t.Run("successful sends", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)

		batch := []education.PersonEvent{{ID: 1}, {ID: 2}, {ID: 3}}

		repo.EXPECT().Lock(gomock.Any()).Times(8).Return(batch, nil)
		repo.EXPECT().Remove(gomock.Any()).Times(3)
		sender.EXPECT().Send(gomock.Any()).Times(24)

		cfg := retranslator.Config{
			ChannelSize: 0,

			ConsumerCount:  2,
			ConsumeSize:    10,
			ConsumeTimeout: 100 * time.Millisecond,

			ProducerCount:  1,
			ProduceTimeout: 100 * time.Millisecond,

			CollectSize:     10,
			CollectMaxDelay: 330 * time.Millisecond,

			WorkerCount: 2,
			WorkTimeout: 100 * time.Millisecond,

			Repo:   repo,
			Sender: sender,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		r := cfg.Start(ctx)
		time.Sleep(390 * time.Millisecond)
		r.Close()
	})

	t.Run("failed sends", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)

		batch := []education.PersonEvent{{ID: 1}, {ID: 2}, {ID: 3}}

		repo.EXPECT().Lock(gomock.Any()).Times(8).Return(batch, nil)
		repo.EXPECT().Unlock(gomock.Any()).Times(3)
		sender.EXPECT().Send(gomock.Any()).Times(24).Return(errors.New("unknown error"))

		cfg := retranslator.Config{
			ChannelSize: 0,

			ConsumerCount:  2,
			ConsumeSize:    10,
			ConsumeTimeout: 100 * time.Millisecond,

			ProducerCount:  3,
			ProduceTimeout: 50 * time.Millisecond,

			CollectSize:     10,
			CollectMaxDelay: 330 * time.Millisecond,

			WorkerCount: 2,
			WorkTimeout: 100 * time.Millisecond,

			Repo:   repo,
			Sender: sender,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		r := cfg.Start(ctx)
		time.Sleep(390 * time.Millisecond)
		r.Close()
	})
}

func TestStart2(t *testing.T) {
	tests := []struct {
		name            string
		repoCfg         DummyRepoCfg
		senderCfg       DummySenderCfg
		retranslatorCfg retranslator.Config
	}{
		{
			"small",
			DummyRepoCfg{
				Size:       20,
				Latency:    25 * time.Millisecond,
				ErrPer100K: 30_000,
			},
			DummySenderCfg{
				Size:       20,
				Latency:    50 * time.Millisecond,
				ErrPer100K: 30_000,
			},
			retranslator.Config{
				ChannelSize:    0,
				ConsumerCount:  3,
				ConsumeSize:    4,
				ConsumeTimeout: 50 * time.Millisecond,

				ProducerCount:  10,
				ProduceTimeout: 100 * time.Millisecond,

				CollectSize:     4,
				CollectMaxDelay: 50 * time.Millisecond,

				WorkerCount: 2,
				WorkTimeout: 100 * time.Millisecond,
			},
		},
		{
			"big",
			DummyRepoCfg{
				Size:       20_000,
				Latency:    25 * time.Millisecond,
				ErrPer100K: 30_000,
			},
			DummySenderCfg{
				Size:       20_000,
				Latency:    50 * time.Millisecond,
				ErrPer100K: 30_000,
			},
			retranslator.Config{
				ChannelSize:    0,
				ConsumerCount:  2,
				ConsumeSize:    100,
				ConsumeTimeout: 100 * time.Millisecond,

				ProducerCount:  400,
				ProduceTimeout: 50 * time.Millisecond,

				CollectSize:     100,
				CollectMaxDelay: 50 * time.Millisecond,

				WorkerCount: 4,
				WorkTimeout: 100 * time.Millisecond,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repoCfg.New()
			sender := tt.senderCfg.New()

			tt.retranslatorCfg.Repo = repo
			tt.retranslatorCfg.Sender = sender

			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()

			tm := time.NewTimer(10 * time.Second)
			tk := time.NewTicker(100 * time.Millisecond)
			r := tt.retranslatorCfg.Start(ctx)

		loop:
			for ctx.Err() == nil {
				select {
				case <-ctx.Done():
					break loop
				case <-tm.C:
					break loop
				case <-tk.C:
					n := sender.Len()
					if tt.repoCfg.Size <= 20 {
						t.Logf("sent events: %d\r", n)
					}
					if n == tt.repoCfg.Size {
						break loop
					}
				}
			}

			tm.Stop()
			tk.Stop()
			r.Close()

			deferred := repo.deferred
			processed := repo.processed.Values()
			removed := repo.removed
			sent := sender.sent

			sortEvents := func(events []uint64) {
				sort.Slice(events, func(i, j int) bool {
					return events[i] < events[j]
				})
			}

			sortEvents(removed)
			sortEvents(sent)

			if tt.repoCfg.Size <= 20 {
				sortEvents(deferred)
				sortEvents(processed)
				t.Logf("deferred : %v\n", deferred)
				t.Logf("processed: %v\n", processed)
				t.Logf("removed  : %v\n", removed)
				t.Logf("sent     : %v\n", sent)
			} else {
				t.Logf("len(deferred) : %v\n", len(deferred))
				t.Logf("len(processed): %v\n", len(processed))
				t.Logf("len(removed)  : %v\n", len(removed))
				t.Logf("len(sent)     : %v\n", len(sent))
			}

			if len(processed) != 0 {
				t.Error("")
			}

			if got := len(deferred) + len(removed) + len(processed); got != tt.repoCfg.Size {
				t.Errorf("the number of repository events has changed. got %d, want %d", got, tt.repoCfg.Size)
			}

			checkEvents := func() bool {
				events := make([]bool, tt.repoCfg.Size+1)

				for _, ev := range deferred {
					if events[ev] {
						return false
					}
				}

				for _, ev := range processed {
					if events[ev] {
						return false
					}
				}

				for _, ev := range removed {
					if events[ev] {
						return false
					}
				}

				return true
			}

			if !checkEvents() {
				t.Error("repository event IDs are not consistent")
			}

			if got := repo.missed; got != 0 {
				t.Errorf("missed repository updates = %d, want 0", got)
			}

			if got := len(processed); got != 0 {
				t.Errorf("there are %d sent left in processing, want 0", got)
			}

			if !reflect.DeepEqual(removed, sent) {
				t.Error("IDs of the removed and sent events do not match")
			}
		})
	}

}
