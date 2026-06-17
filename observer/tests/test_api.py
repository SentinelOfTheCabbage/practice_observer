import pytest

from conftest import Service


def test_basic(service: Service):
    response = service.get()
    assert response.text == "Hi mate! It's main page afaik."
    assert response.status_code == 200


@pytest.mark.parametrize(
    "page, expect_fail",
    [
        (-1, True),
        ("nil", True),
        (None, False),
        (1, False),
    ],
)
def test_page(service: Service, page: int, expect_fail: bool):
    response = service.get(params={"page": page})
    if expect_fail:
        assert response.status_code == 404
    else:
        assert response.status_code == 200
