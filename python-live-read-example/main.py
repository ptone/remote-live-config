import json
import os.path
import time
from settings import Settings

if __name__ == "__main__":

    settings = Settings("/tmp/demo.json")
    try:
        while True:
            print(settings.msg)
            time.sleep(1)
    except KeyboardInterrupt:
        settings.stop()
    settings.join()