import setuptools

setuptools.setup(
    name="grpc-education-person-api",
    version="1.0.0",
    author="rusdevop",
    author_email="rusdevops@gmail.com",
    description="GRPC python client for education-person-api",
    url="https://github.com/aaa2ppp/ozonmp-education-person-api",
    packages=setuptools.find_packages(),
    package_data={"ozonmp.education_person_api.v1": ["education_person_api_pb2.pyi"]},
    python_requires='>=3.5',
)