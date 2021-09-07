from flask import Flask, request
from recipe_scrapers import scrape_me

app = Flask(__name__)

@app.route('/scrape', methods=['POST'])
def scrape():
    return  scrape_me(request.form['url']).instructions() 

if __name__ == '__main__':
   app.run(debug=True, host='0.0.0.0')