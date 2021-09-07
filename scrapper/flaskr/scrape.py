from flask import request, Blueprint
from recipe_scrapers import scrape_me

bp = Blueprint('scrape', __name__, url_prefix='/scrape')

def scrape_url(scraper, url):
    recipe = scraper(url)

    return {
        'title': recipe.title(),
        'time': recipe.total_time(),
        'servings': recipe.yields(),
        'ingredients': recipe.ingredients(),
        'instructions': recipe.instructions().splitlines(),
        'image': recipe.image(),
        'url': url,
    }

@bp.route('/', methods=['POST'])
def scrape():
    return scrape_url(scraper=scrape_me, url=request.form['url']) 
