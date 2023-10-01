package api

import (
	"context"
	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	"github.com/aaa2ppp/ozonmp-education-person-api/internal/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"

	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/aaa2ppp/ozonmp-education-person-api/pkg/education-person-api"
)

var (
	totalPersonNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "education_person_api_person_not_found_total",
		Help: "Total number of persons that were not found",
	})
)

type personAPI struct {
	pb.UnimplementedEducationPersonApiServiceServer
	repo repo.Repo
}

// NewPersonAPI returns api of education-person-api service
func NewPersonAPI(r repo.Repo) pb.EducationPersonApiServiceServer {
	return &personAPI{repo: r}
}

const unimplemented = "Unimplemented"

func timestampToTime(t *timestamppb.Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}
	return t.AsTime()
}

func timeToTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

func (o *personAPI) DescribePersonV1(
	ctx context.Context,
	req *pb.DescribePersonV1Request,
) (*pb.DescribePersonV1Response, error) {
	const op = "DescribePersonV1"

	log.Debug().Msgf("%s - req=%v", op, req)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("%s - invalid argument", op)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	person, err := o.repo.DescribePerson(ctx, req.PersonId)
	if err != nil && err != service.ErrNotFound { // TODO: Здесь мы общаемся с Repo а ошибка прилетает от Service
		log.Error().Err(err).Msgf("%s - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if person == nil {
		log.Debug().Uint64("personId", req.PersonId).Msg("person not found")
		totalPersonNotFound.Inc()

		return nil, status.Error(codes.NotFound, "person not found")
	}

	log.Debug().Msgf("%s - success", op)

	return &pb.DescribePersonV1Response{
		Person: &pb.Person{
			Id:         person.ID,
			FirstName:  person.FirstName,
			MiddleName: person.MiddleName,
			LastName:   person.LastName,
			Birthday:   timeToTimestamp(person.Birthday),
			Sex:        pb.Sex(person.Sex),
			Education:  pb.Education(person.Education),
		},
	}, nil
}

func (o *personAPI) CreatePersonV1(
	ctx context.Context,
	req *pb.CreatePersonV1Request,
) (*pb.CreatePersonV1Response, error) {
	const op = "CreatePersonV1"

	log.Debug().Msgf("%s - req=%v", op, req)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("%s - invalid argument", op)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	p := req.Person

	id, err := o.repo.CreatePerson(ctx, model.Person{
		FirstName:  p.FirstName,
		MiddleName: p.MiddleName,
		LastName:   p.LastName,
		Birthday:   timestampToTime(p.Birthday),
		Sex:        model.Sex(p.Sex),
		Education:  model.Education(p.Education),
	})

	if err != nil {
		log.Error().Err(err).Msgf("%s - can't create person", op)

		return nil, status.Error(codes.Internal, err.Error()) // TODO: нужно ли здесь возвращать сырую ошибку?
	}

	return &pb.CreatePersonV1Response{PersonId: id}, nil
}

func (o *personAPI) ListPersonV1(
	ctx context.Context,
	req *pb.ListPersonV1Request,
) (*pb.ListPersonV1Response, error) {
	const op = "ListPersonV1"

	log.Debug().Msgf("%s - req=%v", op, req)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("%s - invalid argument", op)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	persons, err := o.repo.ListPerson(ctx, req.Cursor, req.Limit)

	if err != nil && err != service.ErrNotFound { // TODO: why *service*.ErrNotFound?
		log.Error().Err(err).Msgf("%s - can't create p", op)

		// TODO: is it really necessary to returns the raw error?
		return nil, status.Error(codes.Internal, err.Error())
	}

	buf := make([]pb.Person, len(persons))
	for i, p := range persons {
		buf[i] = pb.Person{
			Id:         p.ID,
			FirstName:  p.FirstName,
			MiddleName: p.MiddleName,
			LastName:   p.LastName,
			Birthday:   timeToTimestamp(p.Birthday),
			Sex:        pb.Sex(p.Sex),
			Education:  pb.Education(p.Education),
		}
	}

	pp := make([]*pb.Person, len(buf))
	for i := range buf {
		pp[i] = &buf[i]
	}

	return &pb.ListPersonV1Response{
		Persons: pp,
	}, nil
}

func (o *personAPI) RemovePersonV1(
	ctx context.Context,
	req *pb.RemovePersonV1Request,
) (*pb.RemovePersonV1Response, error) {
	const op = "ListPersonV1"

	log.Debug().Msgf("%s - req=%v", op, req)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("%s - invalid argument", op)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ok, err := o.repo.RemovePerson(ctx, req.PersonId)

	if err != nil && err != service.ErrNotFound { // TODO: why *service*.ErrNotFound?
		log.Error().Err(err).Msgf("%s - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RemovePersonV1Response{Ok: ok}, nil

}

func (o *personAPI) UpdatePersonV1(
	ctx context.Context,
	req *pb.UpdatePersonV1Request,
) (*pb.UpdatePersonV1Response, error) {
	const op = "UpdatePersonV1"

	log.Debug().Msgf("%s - req=%v", op, req)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("%s - invalid argument", op)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// TODO: здесь гонка за запись в репо, мы сначала читаем, потом записываем

	person, err := o.repo.DescribePerson(ctx, req.PersonId)
	if err != nil && err != service.ErrNotFound { // TODO: Здесь мы общаемся с Repo а ошибка прилетает от Service
		log.Error().Err(err).Msgf("%s - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if person == nil {
		log.Debug().Uint64("personId", req.PersonId).Msg("person not found")
		totalPersonNotFound.Inc()

		return nil, status.Error(codes.NotFound, "person not found")
	}

	p := req
	if p.FirstName != nil {
		person.FirstName = p.FirstName.Value
	}
	if p.MiddleName != nil {
		person.MiddleName = p.MiddleName.Value
	}
	if p.LastName != nil {
		person.LastName = p.LastName.Value
	}
	if p.Birthday != nil {
		person.Birthday = timestampToTime(p.Birthday.Value)
	}
	if p.Sex != nil {
		person.Sex = model.Sex(p.Sex.Value)
	}
	if p.Education != nil {
		person.Education = model.Education(p.Education.Value)
	}

	err = o.repo.UpdatePerson(ctx, p.PersonId, *person)
	if err != nil && err != service.ErrNotFound { // TODO: Здесь мы общаемся с Repo а ошибка прилетает от Service
		log.Error().Err(err).Msgf("%s - failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if err == service.ErrNotFound {
		log.Debug().Uint64("personId", req.PersonId).Msg("person not found")
		totalPersonNotFound.Inc()

		return nil, status.Error(codes.NotFound, "person not found")
	}

	return &pb.UpdatePersonV1Response{Ok: true}, nil
}
