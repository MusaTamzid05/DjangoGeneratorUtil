from .views import somethng
from .views import somethng
from .views import city
from .views import city
from .views import city
from .views import city
from django.urls import path
from django.urls import include
from .views import home
from .views import detail
from .views import about

urlpatterns = [
        path("", home, name="home"),
        path("post/<int:post_id>", detail ,name="detail"),
        path("about/", about,name="about"),
	path("city/",city),


		path("city/",city),

	path("city/",city),

	path("city/",city),

	path("somethng/",somethng),

	path("somethng/",somethng),

	]
