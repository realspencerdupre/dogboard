import io
import logging
import os
import zipfile
from pathlib import Path

from django.core.files import File
from django.http import FileResponse, HttpResponseBadRequest, JsonResponse
from django.shortcuts import redirect, render

from .forms import SoundForm
from .models import Sound


logger = logging.getLogger(__name__)


def index(request):
    sounds = []
    for sound in Sound.objects.filter(archived=False):
        sounds.append(sound.get_api_obj())
    return render(request, "mainapp/index.html", {"sounds": sounds})


def handle_export(request):
    if request.method == "POST":
        # Get the list of names from the POST request
        names = request.POST.getlist("names")

        # Create a Zip file in memory
        zip_bytes = io.BytesIO()
        zip_file = zipfile.ZipFile(zip_bytes, "w")

        # Iterate through the names and add the corresponding sounds to the Zip file
        for name in names:
            sound = Sound.objects.get(name=name)

            sound_filename = sound.name + ".mp3"
            sound_path = sound.audio.path
            sound_arcname = os.path.join(sound.name, sound_filename)
            zip_file.write(sound_path, arcname=sound_arcname)

            pic_filename = sound.name + ".png"
            pic_path = sound.icon.path
            pic_arcname = os.path.join(sound.name, pic_filename)
            zip_file.write(pic_path, arcname=pic_arcname)

        # Close the Zip file
        zip_file.close()

        # Rewind the in-memory Zip file
        zip_bytes.seek(0)

        # Create the HTTP response with the Zip file as the content
        response = FileResponse(zip_bytes, content_type="application/zip")
        response["Content-Disposition"] = 'attachment; filename="dogboard-export.zip"'
        return response

    return HttpResponseBadRequest()


def handle_delete(request):
    if request.method == "POST":
        # Get the list of names from the POST request
        names = request.POST.getlist("names")

        for name in names:
            sound = Sound.objects.get(name=name)
            sound.delete()

        return redirect("mainapp:index")

    return HttpResponseBadRequest()


def handle_import(request):
    if request.method == "POST":
        zip_handle = zipfile.ZipFile(request.FILES["zip_file"], "r")
        file_list = zip_handle.namelist()
        file_list = sorted(file_list, key=lambda x: x)
        file_list = [p for p in file_list if not p.endswith('/')]
        for audio, icon in zip(file_list[::2], file_list[1::2]):
            name = os.path.dirname(audio)
            logger.info(f'Importing {name}')
            audio_file = File(zip_handle.open(audio))
            icon_file = File(zip_handle.open(icon))
            sound = Sound(
                name=name,
                description="test",
                source="test",
            )
            sound.save()
            sound.audio.save(f"{name}.mp3", audio_file)
            sound.icon.save(f"{name}{Path(icon).suffix}", icon_file)
            sound.load_into_cache()
        zip_handle.close()
        return redirect("mainapp:index")

    return HttpResponseBadRequest()


def add(request):
    if request.method == "GET":
        form = SoundForm()
        return render(request, "mainapp/add.html", {"form": form})
    elif request.method == "POST":
        form = SoundForm(request.POST, request.FILES)
        if form.is_valid():
            sound = form.save()
            sound.load_into_cache()
            return redirect("mainapp:index")
        else:
            return render(request, "mainapp/add.html", {"form": form})
