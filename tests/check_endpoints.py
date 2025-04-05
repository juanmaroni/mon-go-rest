from sys import exit, argv
from config.logger import setup_logger
from config.environment import Environment
from requests_test.requests_test import (
    get_homepage_test, get_pokeapi_all_test, get_pokeapi_by_region_test,
    get_pokeapi_by_id_test, get_pokeapi_by_region_error_test,
    get_pokeapi_by_id_error_test, post_pokeapi_error_test
)

DEV = {
    "host": "localhost",
    "port": 3333
}

PRO = {} # Placeholder


def check_environment_parameters(env: dict):
    return "host" in env and "port" in env and env["host"] and env["port"]


def run_tests(env: Environment):
    base_uri = env.get_base_uri()
    api_uri = env.get_api_base_uri()

    logger.info(f"Testing method GET for URI '{base_uri}'")
    logger.info(get_homepage_test(base_uri))

    logger.info(f"Testing method GET for URI '{api_uri}'")
    logger.info(get_pokeapi_all_test(api_uri))

    test_uri = f"{api_uri}/kanto"
    logger.info(f"Testing method GET for URI '{test_uri}'")
    logger.info(get_pokeapi_by_region_test(test_uri))

    test_uri = f"{api_uri}/151"
    logger.info(f"Testing method GET for URI '{test_uri}'")
    logger.info(get_pokeapi_by_id_test(test_uri))

    test_uri = f"{api_uri}/otnak"
    logger.info(f"Testing method GET for URI '{test_uri}'")
    logger.info(get_pokeapi_by_region_error_test(test_uri))

    test_uri = f"{api_uri}/9999"
    logger.info(f"Testing method GET for URI '{test_uri}'")
    logger.info(get_pokeapi_by_id_error_test(test_uri))

    test_uri = f"{api_uri}/9999"
    logger.info(f"Testing method POST for URI '{test_uri}'")
    logger.info(post_pokeapi_error_test(test_uri))


if __name__ == "__main__":
    # Example: python check_endpoints.py dev
    logger = setup_logger()
    args = argv[1:]
    logger.info(f"List of arguments: {args}")
    n_args = len(args)

    if n_args != 1:
        logger.error("Incorrect number of arguments.")
        exit(1)

    env = args[0].lower()
    
    if env == "dev":
        logger.info(f"Running tests for '{env}' environment.")
        
        if not check_environment_parameters(DEV):
            logger.error("Check 'ENV' dict.")
            exit(1)
        
        env = Environment(DEV["host"], DEV["port"])
        run_tests(env)
    elif env == "pro":
        logger.info(f"Running tests for '{env}' environment.")

        if not check_environment_parameters(PRO):
            logger.error("Check 'PRO' dict.")
            exit(1)
        
        env = Environment(PRO["host"], PRO["port"])
        run_tests(env)
    else:
        logger.error(f"Incorrect argument '{env}'.")
        exit(1)

    logger.info("All tests passed.")
    exit(0)
