package retranslator

import (
	"github.com/aaa2ppp/zonmp-education-person-api/internal/app/workerpool"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/lib/log/lo"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/mocks"
	"github.com/aaa2ppp/zonmp-education-person-api/internal/model/person"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	ctrl := gomock.NewController(t)
	repo := mocks.NewMockEventRepo(ctrl)
	sender := mocks.NewMockEventSender(ctrl)

	repo.EXPECT().Lock(gomock.Any()).Times(1).Return(
		[]person.PersonEvent{{ID: 1}, {ID: 2}},
		nil,
	)
	sender.EXPECT().Send(gomock.Any()).Times(2)
	repo.EXPECT().Remove(gomock.Any()).Times(1)

	cfg := Config{
		ChannelSize:       0,
		ConsumerCount:     1,
		ConsumerBatchSize: 10,
		ConsumerTimeout:   10 * time.Second,
		ProducerCount:     1,
		ProducerTimeout:   10 * time.Second,
		WorkerPoolCfg: workerpool.Config{
			CleanerCount:  1,
			UnlockerCount: 1,
			BatchSize:     10,
			BatchTimeout:  1 * time.Second,
			FailTimeout:   10 * time.Second,
		},
	}

	lo.DebugEnable = true

	r := Start(repo, sender, cfg)
	time.Sleep(2 * time.Second)
	r.Close()
}
