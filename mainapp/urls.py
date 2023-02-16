from django.urls import path

from . import views

app_name = "mainapp"

urlpatterns = [
    path("", views.index, name="index"),
    path("handle-export/", views.handle_export, name="handle-export"),
    path("handle-delete/", views.handle_delete, name="handle-delete"),
    path("handle-import/", views.handle_import, name="handle-import"),
    path("add/", views.add, name="add"),
]
