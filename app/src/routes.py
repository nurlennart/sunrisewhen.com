from flask import render_template, request, make_response, redirect, send_from_directory, jsonify
import pytz
import json
from localizer import localizer
from app import app, limiter


@app.route('/')
def index():
    if not request.cookies.get('times'):
        try:
            local = localizer(request.headers['X-Real-IP'])
            suntimes = local.fitTimezone(local.sunapi())
        except pytz.exceptions.UnknownTimeZoneError:
            return render_template('settings.html', message='something went wrong, bad coordinates perhaps?')
        except TypeError:
            return render_template('settings.html', message='something went wrong, bad coordinates perhaps?')
        except KeyError:
            res = redirect('/', 302)
            res.set_cookie('times', '', max_age=0)
            return(res)

        res = make_response(render_template('index.html', suntimes=suntimes, local=local, request=request))
        res.set_cookie('lat', str(local.lat), max_age=60*60*24*3, secure=True, httponly=False)
        res.set_cookie('lon', str(local.lon), max_age=60*60*24*3, secure=True, httponly=False)
        res.set_cookie('zoneName', str(local.zone), max_age=60*60*24*3, secure=True, httponly=True)
        res.set_cookie('times', json.dumps(suntimes), max_age=60*60*12, secure=True, httponly=False)
        return res
    else:
        suntimes = json.loads(request.cookies.get('times'))
        return render_template('index.html', suntimes=suntimes, request=request)


@app.route('/update')
@limiter.limit("100/minute")
def update():
    try:
        timedict = []
        local = localizer(request.headers['X-Real-IP'])
        suntimes = local.fitTimezone(local.sunapi())
        for time in suntimes:
            timedict.append(time)
        res = jsonify(timedict)
        res.set_cookie('times', json.dumps(suntimes), max_age=60*60*12, secure=True, httponly=False)
        return(res)
    except Exception:
        return('error')


@app.route('/credits')
def credit():
    return render_template('credits.html')


@app.route('/android-chrome-192x192.png')
@app.route('/android-chrome-512x512.png')
@app.route('/browserconfig.xml')
@app.route('/mstile-150x150.png')
@app.route('/robots.txt')
@app.route('/favicon.ico')
def static_from_root():
    return send_from_directory(app.static_folder, request.path[1:])
