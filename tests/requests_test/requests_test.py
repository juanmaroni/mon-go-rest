import requests


COUNT_POKEMON_KANTO_COLL = 151
COUNT_POKEMON_JOHTO_COLL = 100
COUNT_ALL_POKEMON_DB = COUNT_POKEMON_KANTO_COLL + COUNT_POKEMON_JOHTO_COLL
JSON_ID_151 = {
    "id": 151,
    "name": "Mew",
    "type1": "Psychic",
    "type2": "",
}


class ExpectedResponse:
    def __init__(self, status_code, msg, count_json_records=0, json=None):
        self.status_code = status_code
        self.msg = msg
        self.count_json_records = count_json_records
        self.json = json


def get_homepage_test(uri: str):
    expected = ExpectedResponse(
        200,
        "PokéAPI Mini home page."
    )
    response = requests.get(uri)

    compare_responses(response, expected)

    return ("Test 'get_homepage_test' passed.")


def get_pokeapi_all_test(uri: str):
    expected = ExpectedResponse(
        200,
        "",
        COUNT_ALL_POKEMON_DB,
        JSON_ID_151
    )
    response = requests.get(uri)
    
    compare_responses(response, expected)

    return ("Test 'get_pokeapi_all_test' passed.")


def get_pokeapi_by_region_test(uri: str):
    expected = ExpectedResponse(
        200,
        "",
        COUNT_POKEMON_KANTO_COLL,
        JSON_ID_151
    )
    response = requests.get(uri)
    
    compare_responses(response, expected)

    return ("Test 'get_pokeapi_by_region_test' passed.")


def get_pokeapi_by_id_test(uri: str):
    expected = ExpectedResponse(
        200,
        "",
        1,
        JSON_ID_151
    )
    response = requests.get(uri)

    compare_responses(response, expected)

    return ("Test 'get_pokeapi_by_id_test' passed.")


def get_pokeapi_by_region_error_test(uri: str):
    expected = ExpectedResponse(
        404,
        "404 Not Found"
    )
    response = requests.get(uri)

    compare_responses(response, expected)

    return ("Test 'get_pokeapi_by_region_error_test' passed.")


def get_pokeapi_by_id_error_test(uri: str):
    expected = ExpectedResponse(
        404,
        "404 Not Found"
    )
    response = requests.get(uri)

    compare_responses(response, expected)

    return ("Test 'get_pokeapi_by_id_error_test' passed.")


def post_pokeapi_error_test(uri: str):
    expected = ExpectedResponse(
        405,
        "405 Method Not Allowed"
    )
    response = requests.post(uri)

    compare_responses(response, expected)

    return ("Test 'post_pokeapi_error_test' passed.")


def compare_responses(response: requests.Response, expected: ExpectedResponse):
    assert response.status_code == expected.status_code

    if expected.json:
        response_json = response.json()

        if type(response_json) == list:
            assert len(response_json) == expected.count_json_records
            compare_json_record(response_json[150], expected.json)
        else:
            compare_json_record(response_json, expected.json)
    else:
        assert response.content.decode("utf-8") == expected.msg


def compare_json_record(response_json: dict, expected_json: dict):
    assert response_json["id"] == expected_json["id"]
    assert response_json["name"] == expected_json["name"]
    assert response_json["type1"] == expected_json["type1"]
    assert response_json["type2"] == expected_json["type2"]
