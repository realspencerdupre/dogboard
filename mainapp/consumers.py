# chat/consumers.py
import json
import logging

import time

from channels.generic.websocket import WebsocketConsumer

from .models import Sound

LAG_THRES = 1.0


logger = logging.getLogger(__name__)


class ChatConsumer(WebsocketConsumer):
    def connect(self):
        self.accept()

    def disconnect(self, code):
        pass

    def receive(self, text_data, bytes_data=None):
        text_data_json = json.loads(text_data)
        logger.info(f'Received data\n {text_data_json}')

        if "rpc" in text_data_json:
            if text_data_json["rpc"] == "playsound":
                self.playsound_rpc(text_data_json)

    def playsound_rpc(self, _json):
        name = _json["args"]
        sound = Sound.objects.get(name=name)
        lag = time.time() - float(_json["timestamp"])
        if (lag) > LAG_THRES:
            logger.warning(f"Lag time ({lag:.1f}s) too long")
            return
        sound.play_from_cache()
