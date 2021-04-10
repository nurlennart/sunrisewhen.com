from flask import Flask
from flask_limiter import Limiter
from flask_limiter.util import get_ipaddr


app = Flask(__name__)
limiter = Limiter(
    app,
    key_func=get_ipaddr,
    default_limits=['100 per hour']
)

import routes
