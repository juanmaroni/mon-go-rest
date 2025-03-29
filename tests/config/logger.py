from sys import stdout
import logging


def setup_logger():
    logging.basicConfig(
        stream=stdout,
        level=logging.INFO, 
        format='[%(asctime)s] %(levelname)s\t- %(message)s',
    )

    return logging.getLogger(__name__)
