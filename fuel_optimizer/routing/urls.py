from django.urls import path
from .views import RouteView

urlpatterns = [
    path('optimize-route/', RouteView.as_view(), name='optimize-route'),
]

