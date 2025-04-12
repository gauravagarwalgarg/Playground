import random
import string
from flask import Flask, request, redirect, jsonify
from flask_sqlalchemy import SQLAlchemy

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///urls.db"
db = SQLAlchemy(app)

class URLMapping(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    long_url = db.Column(db.String(2048), nullable=False)
    short_code = db.Column(db.String(10), unique=True, nullable=False)


def generate_short_code(length=6):
    characters = string.ascii_letters + string.digits
    return ''.join(random.choices(characters, k=length))

@app.route("/shorten", methods=["POST"])
def shorten_url():
    data = request.get_json()
    long_url = data.get("long_url")

    exisiting_entry = URLMapping.query.filter_by(long_url=long_url).first()
    if exisiting_entry:
        return jsonify({"short_url": f"http://localhost:5000/{exisiting_entry.short_code}"})

    short_code = generate_short_code()
    while URLMapping.query.filter_by(short_code=short_code).first():
        short_code = generate_short_code()

    new_entry = URLMapping(long_url=long_url, short_code=short_code)
    db.session.add(new_entry)
    db.session.commit()

    return jsonify({"short_url": f"http://localhost:5000/{short_code}"})


@app.route("/<short_code>")
def redirect_url(short_code):
    url_entry = URLMapping.query.filter_by(short_code=short_code).first()
    if url_entry:
        return redirect(url_entry.long_url)
    return jsonify({"error": "Short URL not found"}), 404


if __name__ == "__main__":
    with app.app_context():
        db.create_all()
    app.run(debug=True)
