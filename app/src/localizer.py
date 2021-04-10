from flask import request
import requests
import datetime
import config as config
import pytz
import geoip2.database

TIMEZONEDB_URL = 'http://api.timezonedb.com/v2.1/get-time-zone?key={}&by=position&lat={}&lng={}&format=json'
SUNAPI_URL = 'https://api.sunrise-sunset.org/json?{}&formatted=0'


class localizer():
    def __init__(self, ip):
        self.lat = ''
        self.lon = ''
        self.zone = ''
        self.ip = ip

    def location(self):
        if request.cookies.get('lat') and request.cookies.get('lon'):
            self.lat = request.cookies.get('lat')
            self.lon = request.cookies.get('lon')
        else:
            with geoip2.database.Reader('GeoLite2-City.mmdb') as geo_db:
                resp = geo_db.city(self.ip)
                self.lat = resp.location.latitude
                self.lon = resp.location.longitude
        loc_string = f'lat={self.lat}&lng={self.lon}'
        return(loc_string)

    def timezone(self):
        if not request.cookies.get('zoneName'):
            api = requests.get(TIMEZONEDB_URL.format(config.api.key, self.lat, self.lon)).json()
            self.zone = api['zoneName']
        else:
            self.zone = request.cookies.get('zoneName')

    def sunapi(self):
        api = requests.get(SUNAPI_URL.format(self.location())).json()
        return(api)

    def fitTimezone(self, suntimes):
        self.timezone()
        utc_rise = datetime.datetime.strptime(suntimes['results']['sunrise'], '%Y-%m-%dT%H:%M:%S%z')
        utc_set = datetime.datetime.strptime(suntimes['results']['sunset'], '%Y-%m-%dT%H:%M:%S%z')
        local_rise = utc_rise.astimezone(pytz.timezone(self.zone)).strftime('%I:%M:%S %p')
        local_set = utc_set.astimezone(pytz.timezone(self.zone)).strftime('%I:%M:%S %p')
        final = [local_rise, local_set, utc_rise.strftime('%I:%M:%S %p'), utc_set.strftime('%I:%M:%S %p')]
        return(final)
