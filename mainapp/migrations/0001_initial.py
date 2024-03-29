# Generated by Django 4.1.4 on 2023-01-09 00:41

from django.db import migrations, models

import mainapp.models


class Migration(migrations.Migration):

    initial = True

    dependencies = []

    operations = [
        migrations.CreateModel(
            name="Sound",
            fields=[
                (
                    "id",
                    models.BigAutoField(
                        auto_created=True,
                        primary_key=True,
                        serialize=False,
                        verbose_name="ID",
                    ),
                ),
                ("name", models.CharField(max_length=128)),
                ("description", models.CharField(max_length=256)),
                ("source", models.CharField(max_length=256)),
                ("audio", models.FileField(upload_to=mainapp.models.sound_upload_to)),
                ("icon", models.ImageField(upload_to=mainapp.models.sound_upload_to)),
                ("created", models.DateTimeField(auto_now_add=True)),
                ("active", models.BooleanField(default=True)),
            ],
        ),
    ]
