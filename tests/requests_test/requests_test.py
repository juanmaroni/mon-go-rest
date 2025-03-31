import requests


COUNT_POKEMON_KANTO_COLL = 151
COUNT_POKEMON_JOHTO_COLL = 100
COUNT_ALL_POKEMON_DB = COUNT_POKEMON_KANTO_COLL + COUNT_POKEMON_JOHTO_COLL


class ExpectedResponse:
    def __init__(self, status_code, msg, number_records=0):
        self.status_code = status_code
        self.msg = msg
        self.number_records = number_records


def get_homepage(uri: str):
    expected = ExpectedResponse(
        #"http://localhost:3333"
        200,
        "PokéAPI Mini home page."
    )
    response = requests.get(uri)

    compare_message(uri, response, expected)

    return ("\'get_homepage\' test passed.")


def get_pokeapi_all(uri: str):
    expected = ExpectedResponse(
        "http://localhost:3333/api/v1/pokemon",
        200,
        "",
        COUNT_ALL_POKEMON_DB
    )


def get_pokeapi_by_region(uri: str):
    expected = ExpectedResponse(
        #"http://localhost:3333/api/v1/pokemon/kanto",
        200,
        "",
        COUNT_POKEMON_KANTO_COLL
    )


def get_pokeapi_by_id(uri: str):
    expected = ExpectedResponse(
        #"http://localhost:3333/api/v1/pokemon/151",
        200,
        "",
        1
    )


def get_pokeapi_by_region_error(uri: str):
    expected = ExpectedResponse(
        #"http://localhost:3333/api/v1/pokemon/otnak",
        404,
        "404 Not Found"
    )


def get_pokeapi_by_id_error(uri: str):
    expected = ExpectedResponse(
        #"http://localhost:3333/api/v1/pokemon/9999",
        404,
        "404 Not Found"
    )


def post_pokeapi_error(uri: str):
    expected = ExpectedResponse(
        #"http://localhost:3333/api/v1/pokemon/9999",
        405,
        "405 Method Not Allowed"
    )


def compare_message(response: requests.Response, expected: ExpectedResponse):
    assert response.status_code == expected.status_code
    assert response.content.decode("utf-8") == expected.msg


def compare_json(response: requests.Response, expected: ExpectedResponse):
    assert response.status_code == expected.status_code
    assert response.content.decode("utf-8") == expected.msg
