from flask import Flask, request, jsonify, redirect
from flask_sqlalchemy import SQLAlchemy
import random
import string

app = Flask(__name__)

# Database configuration
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///urls.db"
app.config["SQLALCHEMY_TRACK_MODIFICATIONS"] = False

db = SQLAlchemy(app)

# URL mapping model
class URLMapping(db.Model):
    id = db.Column(db.Integer, primary_key=True)
    long_url = db.Column(db.String(2048), nullable=False)
    short_code = db.Column(db.String(6), unique=True, nullable=False)

# Function to generate a short code
def generate_short_code(length=6):
    characters = string.ascii_letters + string.digits
    return "".join(random.choices(characters, k=length))

# Route to shorten a URL
@app.route("/shorten", methods=["POST"])
def shorten_url():
    data = request.get_json()

    # Validate input
    if not data or "url" not in data:
        return jsonify({"error": "Missing 'url' in request body"}), 400
    
    long_url = data.get("url")
    
    if not long_url:
        return jsonify({"error": "URL cannot be empty"}), 400

    # Check if the URL is already shortened
    existing_entry = URLMapping.query.filter_by(long_url=long_url).first()
    if existing_entry:
        return jsonify({"short_url": f"http://127.0.0.1:5000/{existing_entry.short_code}"})
    
    # Generate a new short code
    short_code = generate_short_code()
    
    # Ensure uniqueness of short code
    while URLMapping.query.filter_by(short_code=short_code).first():
        short_code = generate_short_code()

    # Store in database
    new_entry = URLMapping(long_url=long_url, short_code=short_code)
    db.session.add(new_entry)
    db.session.commit()

    return jsonify({"short_url": f"http://127.0.0.1:5000/{short_code}"})

# Route to expand short URL
@app.route("/<short_code>")
def redirect_to_long_url(short_code):
    entry = URLMapping.query.filter_by(short_code=short_code).first()
    if entry:
        return redirect(entry.long_url)
    return jsonify({"error": "Short URL not found"}), 404

# Initialize database
with app.app_context():
    db.create_all()

if __name__ == "__main__":
    app.run(debug=True)

