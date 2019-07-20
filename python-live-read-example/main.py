import json
import os.path
import time

from bunch import Bunch
from watchdog.events import PatternMatchingEventHandler
from watchdog.observers import Observer

class SettingsChangeHandler(PatternMatchingEventHandler):
    def __init__(self, *args, **kwargs):
        self.settings = kwargs.get("settings")
        del(kwargs["settings"])
        super().__init__(*args, **kwargs)
        for path in kwargs.get("patterns"):
            self.update(path)
        
    def update(self, path):
        with open(path) as conf:
            self.settings.update(json.load(conf))

    def on_modified(self, event):
        print("Settings modified on disk, reloading")
        self.update(event.src_path)
    
class Settings(Bunch):
    def __init__(self, settings_path):
        super().__init__()
        event_handler = SettingsChangeHandler(settings=self, patterns=[settings_path])
        self.observer = Observer()
        self.observer.schedule(event_handler, os.path.abspath(os.path.join(settings_path, os.pardir)), recursive=False)
        self.observer.start()
    
    def stop(self):
        self.observer.stop()

    def join(self):
        self.observer.join()



if __name__ == "__main__":

    settings = Settings("/tmp/demo.json")
    try:
        while True:
            print(settings.msg)
            time.sleep(1)
    except KeyboardInterrupt:
        settings.stop()
    settings.join()