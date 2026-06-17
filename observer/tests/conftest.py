import pytest
import os

from dotenv import load_dotenv
import requests


@pytest.fixture(scope="session", autouse=True)
def load_env():
    load_dotenv()


class Service:
    def __init__(self):
        self.__url = "http://localhost"
        self.__port = os.getenv("DOCKER_PORT")

    def request(
        self,
        method: str,
        uri: str,
        params: dict = dict(),
        body: dict = dict(),
        **kwargs,
    ):
        if kwargs:
            body.update(kwargs)
        url = f"{self.__url}:{self.__port}/{uri}"
        return requests.request(method, url, params=params, data=body)

    def get(self, uri: str = "", params: dict = dict(), body: dict = dict(), **kwargs):
        return self.request("GET", uri, params, body, **kwargs)

    def post(self, uri: str = "", params: dict = dict(), body: dict = dict(), **kwargs):
        return self.request("POST", uri, params, body, **kwargs)


@pytest.fixture(name="service")
def _service():
    return Service()
