# Generated by Django 4.1.4 on 2023-01-09 22:59

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ("mainapp", "0001_initial"),
    ]

    operations = [
        migrations.RemoveField(
            model_name="sound",
            name="active",
        ),
        migrations.AddField(
            model_name="sound",
            name="archived",
            field=models.BooleanField(default=False),
        ),
    ]
