from sys import stdout
import logging


def setup_logger():
    logging.basicConfig(
        stream=stdout,
        level=logging.INFO, 
        format='[%(asctime)s] %(levelname)s\t- %(message)s',
    )
    logging.addLevelName(logging.WARNING, "WARN")

    return logging.getLogger(__name__)
