package retranslator

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/loader"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/worker"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

type event = education.PersonEvent

// TODO: remove these dependencies, used only in dummy_repo.go
const (
	created   = education.Created
	processed = education.Processed
)

type (
	workerConfig = worker.Config
	workerJob    = worker.Job
)

var (
	start      = loader.Start
	startN     = loader.StartN
	groupStart = loader.GroupStart
)
