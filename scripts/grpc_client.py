import asyncio

from grpclib.client import Channel

from ozonmp.education_person_api.v1.education_person_api_grpc import EducationPersonApiServiceStub
from ozonmp.education_person_api.v1.education_person_api_pb2 import DescribePersonV1Request

async def main():
    async with Channel('127.0.0.1', 8082) as channel:
        client = EducationPersonApiServiceStub(channel)

        req = DescribePersonV1Request(template_id=1)
        reply = await client.DescribePersonV1(req)
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())
