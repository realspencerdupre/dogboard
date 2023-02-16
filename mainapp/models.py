import logging
import os
import shutil
import time
from pathlib import Path

from django.conf import settings
from django.core.cache import cache
from django.db import models
from django.utils.text import slugify
from pydub import AudioSegment, playback


logger = logging.getLogger(__name__)


def sound_upload_to(instance, filename):
    ext = Path(filename).suffix
    savepath = Path("sounds", instance.name, f"{instance.name}{ext}")

    if Path(settings.MEDIA_ROOT, savepath).exists():
        raise FileExistsError(str(savepath))
    return str(savepath)


class Sound(models.Model):
    name = models.CharField(max_length=128, unique=True)
    description = models.CharField(max_length=256)
    source = models.CharField(max_length=256)
    audio = models.FileField(upload_to=sound_upload_to)
    icon = models.ImageField(upload_to=sound_upload_to)
    created = models.DateTimeField(auto_now_add=True)
    archived = models.BooleanField(default=False)

    def save(self, *args, **kwargs):
        self.name = slugify(self.name)
        super().save(*args, **kwargs)

    def delete(self, *args, **kwargs):
        del_path = os.path.dirname(self.audio.path)
        logger.info(f'Deleting {self.name}')
        shutil.rmtree(del_path, ignore_errors=False)
        super().delete(*args, **kwargs)

    def __str__(self):
        return self.name

    def cache_key(self):
        return f"{settings.SOUND_CACHE_PREFIX}{self.name}"

    def audiosegment(self):
        if self.audio:
            return AudioSegment.from_mp3(self.audio.path)
        return None

    def load_into_cache(self):
        audseg = self.audiosegment()
        if audseg:
            audseg = audseg.normalize()
        cache.set(self.cache_key(), audseg, None)

    def play_from_cache(self):
        playback.play(cache.get(self.cache_key()))

    def get_api_obj(self):
        return {
            'name': self.name,
            'mp3': self.audio.url,
            'png': self.icon.url,
            'playing': False,
        }

    @classmethod
    def load_active_to_cache(cls):
        qs = cls.objects.filter(archived=False)
        for sound in qs:
            sound.load_into_cache()


start_time = time.time()
logger.info('Loading sounds...')
Sound.load_active_to_cache()
logger.info(f'Done loading sounds in {time.time() - start_time:.1f} secs')
