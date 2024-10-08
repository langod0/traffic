"""
URL configuration for djangoProject project.

The `urlpatterns` list routes URLs to views. For more information please see:
    https://docs.djangoproject.com/en/5.1/topics/http/urls/
Examples:
Function views
    1. Add an import:  from my_app import views
    2. Add a URL to urlpatterns:  path('', views.home, name='home')
Class-based views
    1. Add an import:  from other_app.views import Home
    2. Add a URL to urlpatterns:  path('', Home.as_view(), name='home')
Including another URLconf
    1. Import the include() function: from django.urls import include, path
    2. Add a URL to urlpatterns:  path('blog/', include('blog.urls'))
"""


from rest_framework import routers



from newpy.Show_data import Show_data
from newpy.predata import predata
router = routers.DefaultRouter()
# 第一个参数为url前缀，第二个参数是前缀对应的试图集，第三个参数是视图基本名

router.register('show_api', Show_data, basename='show_api')
router.register('pre_api', predata, basename='pre_api')
urlpatterns = [

]

urlpatterns += router.urls

