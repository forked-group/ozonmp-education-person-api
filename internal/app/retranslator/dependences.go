package retranslator

import (
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/loader"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/app/worker"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
)

// TODO: remove this dependence
type event = education.PersonEvent

// TODO: remove these dependencies (used in dummy_repo.go)
const eventTypeCreated = education.Created
const eventStatusProcessed = education.Processed

type workerConfig = worker.Config
type workerJob = worker.Job

var start = loader.Start
var groupStart = loader.GroupStart
var startN = loader.StartN
