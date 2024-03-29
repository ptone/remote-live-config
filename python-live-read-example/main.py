# Copyright 2019 Google Inc. All rights reserved.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to writing, software distributed
# under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
# CONDITIONS OF ANY KIND, either express or implied.
#
# See the License for the specific language governing permissions and
# limitations under the License.

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