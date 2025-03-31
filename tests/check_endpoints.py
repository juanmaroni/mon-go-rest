from sys import exit, argv
from config.logger import setup_logger
from config.environment import Environment
from requests_test.requests_test import get_homepage


DEV = {
    "host": "localhost",
    "port": 3333
}

PRO = {} # Placeholder


def check_environment_parameters(env: dict):
    return "host" in env and "port" in env and env["host"] and env["port"]


if __name__ == "__main__":
    # Example: python requests_test.py dev
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
        base_uri = env.get_base_uri()
        api_uri = env.get_api_base_uri()

        # TODO: Call tests
        
    elif env == "pro":
        logger.info(f"Running tests for '{env}' environment.")

        if not check_environment_parameters(PRO):
            logger.error("Check 'PRO' dict.")
            exit(1)
        
        env = Environment(PRO["host"], PRO["port"])
        base_uri = env.get_base_uri()
        api_uri = env.get_api_base_uri()

        # TODO: Call tests
        
    else:
        logger.error(f"Incorrect argument '{env}'.")
        exit(1)

    logger.info("All tests passed.")
    exit(0)
