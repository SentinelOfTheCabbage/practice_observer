import pytest
import os

@pytest.fixture('service')
class Service:
    def __init__(self):
        self.__url = 'http://localhost'
        self.__port = os.getenv('DOCKER_PORT')

