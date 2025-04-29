from django.urls import path
from . import views

app_name = 'api'  # Add namespace

urlpatterns = [
    path('hello/', views.hello, name='hello'),
]
