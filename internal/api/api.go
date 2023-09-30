package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aaa2ppp/ozonmp-education-person-api/internal/repo"

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

func (o *personAPI) DescribePersonV1(
	ctx context.Context,
	req *pb.DescribePersonV1Request,
) (*pb.DescribePersonV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribePersonV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	person, err := o.repo.DescribePerson(ctx, req.PersonId)
	if err != nil {
		log.Error().Err(err).Msg("DescribePersonV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if person == nil {
		log.Debug().Uint64("personId", req.PersonId).Msg("person not found")
		totalPersonNotFound.Inc()

		return nil, status.Error(codes.NotFound, "person not found")
	}

	log.Debug().Msg("DescribePersonV1 - success")

	return &pb.DescribePersonV1Response{
		Value: &pb.Person{
			Id:  person.ID,
			Foo: person.Foo,
		},
	}, nil
}
