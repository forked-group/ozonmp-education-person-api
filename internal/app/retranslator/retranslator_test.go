package retranslator_test

import (
	"context"
	"errors"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/retranslator"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/mocks"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"github.com/golang/mock/gomock"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	//lo.DebugEnable = true

	t.Run("successful sends", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)

		batch := []person.PersonEvent{{ID: 1}, {ID: 2}, {ID: 3}}

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

			CollectSize:    10,
			CollectTimeout: 330 * time.Millisecond,

			WorkerCount: 2,
			WorkTimeout: 100 * time.Millisecond,

			Repo:   repo,
			Sender: sender,
		}

		r := retranslator.Start(&cfg)
		time.Sleep(390 * time.Millisecond)
		r.Close()
	})

	t.Run("failed sends", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mocks.NewMockEventRepo(ctrl)
		sender := mocks.NewMockEventSender(ctrl)

		batch := []person.PersonEvent{{ID: 1}, {ID: 2}, {ID: 3}}

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

			CollectSize:    10,
			CollectTimeout: 330 * time.Millisecond,

			WorkerCount: 2,
			WorkTimeout: 100 * time.Millisecond,

			Repo:   repo,
			Sender: sender,
		}

		r := retranslator.Start(&cfg)
		time.Sleep(390 * time.Millisecond)
		r.Close()
	})
}

func TestStart2(t *testing.T) {
	//lo.DebugEnable = true

	repoCfg := DummyRepoCfg{
		Size:       20_000,
		Timeout:    100 * time.Millisecond,
		ErrPer100K: 50_000,
	}

	senderCfg := DummySenderCfg{
		Size:       repoCfg.Size,
		Timeout:    50 * time.Millisecond,
		ErrPer100K: 50_000,
	}

	repo := repoCfg.New()
	sender := senderCfg.New()

	cfg := retranslator.Config{
		ChannelSize: 0,

		ConsumerCount:  2,
		ConsumeSize:    100,
		ConsumeTimeout: 100 * time.Millisecond,

		ProducerCount:  100,
		ProduceTimeout: 100 * time.Millisecond,

		CollectSize:    100,
		CollectTimeout: 100 * time.Millisecond,

		WorkerCount: 8,
		WorkTimeout: 100 * time.Millisecond,

		Repo:   repo,
		Sender: sender,
	}

	r := retranslator.Start(&cfg)

	time.Sleep(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r.CloseCtx(ctx)

	sortEvents := func(events []uint64) {
		sort.Slice(events, func(i, j int) bool {
			return events[i] < events[j]
		})
	}

	sortEvents(repo.Removed)
	sortEvents(sender.Events)

	if repoCfg.Size <= 20 {
		t.Logf("repo.Deferred: %v\n", repo.Deferred)
		t.Logf("repo.Removed : %v\n", repo.Removed)
		t.Logf("sender.Events: %v\n", sender.Events)
	} else {
		t.Logf("len(repo.Deferred): %v\n", len(repo.Deferred))
		t.Logf("len(repo.Removed) : %v\n", len(repo.Removed))
		t.Logf("len(sender.Events): %v\n", len(sender.Events))
	}

	if !reflect.DeepEqual(repo.Removed, sender.Events) {
		t.Error("lists of removed events and sent events are not equal")
	}
}
