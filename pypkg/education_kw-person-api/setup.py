import setuptools

setuptools.setup(
    name="grpc-education_kw-person-api",
    version="1.0.0",
    author="rusdevop",
    author_email="rusdevops@gmail.com",
    description="GRPC python client for education_kw-person-api",
    url="https://github.com/aaa2ppp/ozonmp-education-kw-person-api",
    packages=setuptools.find_packages(),
    package_data={"ozonmp.education_kw_person_api.v1": ["education_kw_person_api_pb2.pyi"]},
    python_requires='>=3.5',
)