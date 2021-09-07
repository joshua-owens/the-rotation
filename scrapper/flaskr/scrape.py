from flask import request, Blueprint
from recipe_scrapers import scrape_me

bp = Blueprint('scrape', __name__, url_prefix='/scrape')

@bp.route('/', methods=['POST'])
def scrape():
    return  scrape_me(request.form['url']).instructions() 
