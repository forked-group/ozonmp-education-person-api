package api

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aaa2ppp/ozonmp-education-kw-person-api/internal/repo"

	pb "github.com/aaa2ppp/ozonmp-education-kw-person-api/pkg/education_kw-person-api"
)

var (
	totalTemplateNotFound = promauto.NewCounter(prometheus.CounterOpts{
		Name: "education_kw_person_api_person_not_found_total",
		Help: "Total number of persons that were not found",
	})
)

type personAPI struct {
	pb.UnimplementedOmpTemplateApiServiceServer
	repo repo.Repo
}

// NewTemplateAPI returns api of education_kw-person-api service
func NewTemplateAPI(r repo.Repo) pb.OmpTemplateApiServiceServer {
	return &personAPI{repo: r}
}

func (o *personAPI) DescribeTemplateV1(
	ctx context.Context,
	req *pb.DescribeTemplateV1Request,
) (*pb.DescribeTemplateV1Response, error) {

	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 - invalid argument")

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	person, err := o.repo.DescribeTemplate(ctx, req.TemplateId)
	if err != nil {
		log.Error().Err(err).Msg("DescribeTemplateV1 -- failed")

		return nil, status.Error(codes.Internal, err.Error())
	}

	if person == nil {
		log.Debug().Uint64("personId", req.TemplateId).Msg("person not found")
		totalTemplateNotFound.Inc()

		return nil, status.Error(codes.NotFound, "person not found")
	}

	log.Debug().Msg("DescribeTemplateV1 - success")

	return &pb.DescribeTemplateV1Response{
		Value: &pb.Template{
			Id:  person.ID,
			Foo: person.Foo,
		},
	}, nil
}
