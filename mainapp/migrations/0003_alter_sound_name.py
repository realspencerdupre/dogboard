# Generated by Django 4.1.4 on 2023-02-05 16:02

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("mainapp", "0002_remove_sound_active_sound_archived"),
    ]

    operations = [
        migrations.AlterField(
            model_name="sound",
            name="name",
            field=models.CharField(max_length=128, unique=True),
        ),
    ]
