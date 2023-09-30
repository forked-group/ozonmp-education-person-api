# Generated by the Protocol Buffers compiler. DO NOT EDIT!
# source: ozonmp/education_person_api/v1/education_person_api.proto
# plugin: grpclib.plugin.main
import abc
import typing

import grpclib.const
import grpclib.client
if typing.TYPE_CHECKING:
    import grpclib.server

import validate.validate_pb2
import google.api.annotations_pb2
import google.protobuf.timestamp_pb2
import ozonmp.education_person_api.v1.education_person_api_pb2


class EducationPersonApiServiceBase(abc.ABC):

    @abc.abstractmethod
    async def DescribePersonV1(self, stream: 'grpclib.server.Stream[ozonmp.education_person_api.v1.education_person_api_pb2.DescribePersonV1Request, ozonmp.education_person_api.v1.education_person_api_pb2.DescribePersonV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def CreatePersonV1(self, stream: 'grpclib.server.Stream[ozonmp.education_person_api.v1.education_person_api_pb2.CreatePersonV1Request, ozonmp.education_person_api.v1.education_person_api_pb2.CreatePersonV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def ListPersonV1(self, stream: 'grpclib.server.Stream[ozonmp.education_person_api.v1.education_person_api_pb2.ListPersonV1Request, ozonmp.education_person_api.v1.education_person_api_pb2.ListPersonV1Response]') -> None:
        pass

    @abc.abstractmethod
    async def RemovePersonV1(self, stream: 'grpclib.server.Stream[ozonmp.education_person_api.v1.education_person_api_pb2.RemovePersonV1Request, ozonmp.education_person_api.v1.education_person_api_pb2.RemovePersonV1Response]') -> None:
        pass

    def __mapping__(self) -> typing.Dict[str, grpclib.const.Handler]:
        return {
            '/ozonmp.education_person_api.v1.EducationPersonApiService/DescribePersonV1': grpclib.const.Handler(
                self.DescribePersonV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.education_person_api.v1.education_person_api_pb2.DescribePersonV1Request,
                ozonmp.education_person_api.v1.education_person_api_pb2.DescribePersonV1Response,
            ),
            '/ozonmp.education_person_api.v1.EducationPersonApiService/CreatePersonV1': grpclib.const.Handler(
                self.CreatePersonV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.education_person_api.v1.education_person_api_pb2.CreatePersonV1Request,
                ozonmp.education_person_api.v1.education_person_api_pb2.CreatePersonV1Response,
            ),
            '/ozonmp.education_person_api.v1.EducationPersonApiService/ListPersonV1': grpclib.const.Handler(
                self.ListPersonV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.education_person_api.v1.education_person_api_pb2.ListPersonV1Request,
                ozonmp.education_person_api.v1.education_person_api_pb2.ListPersonV1Response,
            ),
            '/ozonmp.education_person_api.v1.EducationPersonApiService/RemovePersonV1': grpclib.const.Handler(
                self.RemovePersonV1,
                grpclib.const.Cardinality.UNARY_UNARY,
                ozonmp.education_person_api.v1.education_person_api_pb2.RemovePersonV1Request,
                ozonmp.education_person_api.v1.education_person_api_pb2.RemovePersonV1Response,
            ),
        }


class EducationPersonApiServiceStub:

    def __init__(self, channel: grpclib.client.Channel) -> None:
        self.DescribePersonV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.education_person_api.v1.EducationPersonApiService/DescribePersonV1',
            ozonmp.education_person_api.v1.education_person_api_pb2.DescribePersonV1Request,
            ozonmp.education_person_api.v1.education_person_api_pb2.DescribePersonV1Response,
        )
        self.CreatePersonV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.education_person_api.v1.EducationPersonApiService/CreatePersonV1',
            ozonmp.education_person_api.v1.education_person_api_pb2.CreatePersonV1Request,
            ozonmp.education_person_api.v1.education_person_api_pb2.CreatePersonV1Response,
        )
        self.ListPersonV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.education_person_api.v1.EducationPersonApiService/ListPersonV1',
            ozonmp.education_person_api.v1.education_person_api_pb2.ListPersonV1Request,
            ozonmp.education_person_api.v1.education_person_api_pb2.ListPersonV1Response,
        )
        self.RemovePersonV1 = grpclib.client.UnaryUnaryMethod(
            channel,
            '/ozonmp.education_person_api.v1.EducationPersonApiService/RemovePersonV1',
            ozonmp.education_person_api.v1.education_person_api_pb2.RemovePersonV1Request,
            ozonmp.education_person_api.v1.education_person_api_pb2.RemovePersonV1Response,
        )