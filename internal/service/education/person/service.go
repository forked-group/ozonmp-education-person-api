package person

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/genproto/googleapis/type/date"

	model "github.com/aaa2ppp/ozonmp-education-person-api/internal/model/education"
	pb "github.com/aaa2ppp/ozonmp-education-person-api/pkg/education-person-api"
)

type Service struct {
	client pb.EducationPersonApiServiceClient
}

func NewService(client pb.EducationPersonApiServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (p Service) Describe(personID uint64) (*model.Person, error) {
	const op = "Service.Describe"

	ctx := context.TODO()
	req := pb.DescribePersonV1Request{
		PersonId: personID,
	}

	resp, err := p.client.DescribePersonV1(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return pbPersonToModelPerson(resp.Person), nil
}

func (p Service) List(cursor uint64, limit uint64) ([]model.Person, error) {
	const op = "Service.List"

	ctx := context.TODO()
	req := pb.ListPersonV1Request{
		Cursor: cursor,
		Limit:  limit,
	}

	resp, err := p.client.ListPersonV1(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	list := make([]model.Person, 0, len(resp.Persons))

	for _, pbPerson := range resp.Persons {
		person := pbPersonToModelPerson(pbPerson)
		list = append(list, *person)
	}

	return list, nil
}

func (p Service) Create(person model.Person) (uint64, error) {
	const op = "Service.Create"

	ctx := context.TODO()
	req := pb.CreatePersonV1Request{
		Person: modelPersonToPbPerson(person),
	}

	resp, err := p.client.CreatePersonV1(ctx, &req)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return resp.PersonId, nil
}

func (p Service) Update(personID uint64, person model.Person, fields model.PersonField) (bool, error) {
	const op = "Service.Update"

	ctx := context.TODO()
	req := pb.UpdatePersonV1Request{
		PersonId: personID,
		Person:   modelPersonToPbPersonUpdate(person, fields),
	}

	resp, err := p.client.UpdatePersonV1(ctx, &req)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.Ok, nil
}

func (p Service) Remove(personID uint64) (bool, error) {
	const op = "Service.Remove"

	ctx := context.TODO()
	req := pb.RemovePersonV1Request{
		PersonId: personID,
	}

	resp, err := p.client.RemovePersonV1(ctx, &req)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.Ok, nil
}

func pbPersonToModelPerson(pbPerson *pb.Person) *model.Person {
	return &model.Person{
		ID:         pbPerson.Id,
		FirstName:  pbPerson.FirstName,
		MiddleName: pbPerson.MiddleName,
		LastName:   pbPerson.LastName,
		Birthday:   pbDateToModelDate(pbPerson.Birthday),
		Sex:        model.Sex(pbPerson.Sex),
		Education:  model.Education(pbPerson.Education),
	}
}

func pbDateToModelDate(d *date.Date) *model.Date {
	if d == nil {
		return nil
	}

	return model.NewDate(
		int(d.Year),
		time.Month(d.Month),
		int(d.Day),
	)
}

func modelPersonToPbPerson(person model.Person) *pb.Person {
	return &pb.Person{
		FirstName:  person.FirstName,
		MiddleName: person.MiddleName,
		LastName:   person.LastName,
		Birthday:   modelDateToPbDate(person.Birthday),
		Sex:        pb.Sex(person.Sex),
		Education:  pb.Education(person.Education),
	}
}

func modelPersonToPbPersonUpdate(person model.Person, fields model.PersonField) *pb.PersonUpdate {
	var personUpdate pb.PersonUpdate

	if fields.Includes(model.PersonFirstName) {
		personUpdate.FirstName = &pb.StringValue{Value: person.FirstName}
	}

	if fields.Includes(model.PersonMiddleName) {
		personUpdate.MiddleName = &pb.StringValue{Value: person.MiddleName}
	}

	if fields.Includes(model.PersonLastName) {
		personUpdate.LastName = &pb.StringValue{Value: person.LastName}
	}

	if fields.Includes(model.PersonBirthday) {
		personUpdate.Birthday = &pb.DateValue{Value: modelDateToPbDate(person.Birthday)}
	}

	if fields.Includes(model.PersonSex) {
		personUpdate.Sex = &pb.SexValue{Value: pb.Sex(person.Sex)}
	}

	if fields.Includes(model.PersonEducation) {
		personUpdate.Education = &pb.EducationValue{Value: pb.Education(person.Education)}
	}

	return &personUpdate
}

func modelDateToPbDate(d *model.Date) *date.Date {
	if d == nil {
		return nil
	}

	return &date.Date{
		Year:  int32(d.Year()),
		Month: int32(d.Month()),
		Day:   int32(d.Day()),
	}
}
