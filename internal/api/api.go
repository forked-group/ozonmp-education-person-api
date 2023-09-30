package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

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

func (o *personAPI) DescribePersonV1(
	ctx context.Context,
	req *pb.DescribePersonV1Request,
) (*pb.DescribePersonV1Response, error) {
	const op = "DescribePersonV1"

	log.Debug().Msgf("%s: req=%v", op, req)

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("%s - invalid argument", op)

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return nil, status.Error(codes.Unimplemented, unimplemented)

	//person, err := o.repo.DescribePerson(ctx, req.PersonId)
	//if err != nil {
	//	log.Error().Err(err).Msg("DescribePersonV1 -- failed")
	//
	//	return nil, status.Error(codes.Internal, err.Error())
	//}
	//
	//if person == nil {
	//	log.Debug().Uint64("personId", req.PersonId).Msg("person not found")
	//	totalPersonNotFound.Inc()
	//
	//	return nil, status.Error(codes.NotFound, "person not found")
	//}
	//
	//log.Debug().Msg("DescribePersonV1 - success")
	//
	//return &pb.DescribePersonV1Response{
	//	Person: &pb.Person{
	//		Id:         person.ID,
	//		FistName:   person.FistName,
	//		MiddleName: person.MiddleName,
	//		LastName:   person.LastName,
	//		Birthday:   timestamppb.New(person.Birthday),
	//		Sex:        pb.Sex(person.Sex),
	//		Education:  pb.Education(person.Education),
	//	},
	//}, nil
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

	return nil, status.Error(codes.Unimplemented, unimplemented)
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

	return nil, status.Error(codes.Unimplemented, unimplemented)
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

	return nil, status.Error(codes.Unimplemented, unimplemented)
}
