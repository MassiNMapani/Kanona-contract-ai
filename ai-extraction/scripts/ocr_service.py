from flask import Flask, request, jsonify
import fitz  # PyMuPDF
import spacy
import re
from pathlib import Path

app = Flask(__name__)
nlp = spacy.load("en_core_web_sm")

# ✅ Hardcoded, working path to your real uploads folder
UPLOAD_FOLDER = Path("/Users/massinachombemapani/Kanona-contract-ai/backend/uploads")
print("📁 Upload folder resolved to:", UPLOAD_FOLDER)

def extract_text_from_pdf(path):
    text = ""
    doc = fitz.open(str(path))
    for page in doc:
        text += page.get_text()
    return text.strip(), doc.page_count

@app.route('/analyze', methods=['POST'])
def analyze_contract():
    filename = request.json.get("filename")
    if not filename:
        return jsonify({"error": "Filename is required"}), 400

    file_path = UPLOAD_FOLDER / filename
    print("📄 Looking for file at:", file_path)

    if not file_path.exists():
        return jsonify({"error": f"File not found at {file_path}"}), 404

    text, pages = extract_text_from_pdf(file_path)
    doc = nlp(text)

    entities = [{"label": ent.label_, "text": ent.text} for ent in doc.ents]
    tariffs = re.findall(r"(ZMW|USD)?\s?[\d,]+(?:\.\d{2})?", text)

    renegotiation_match = re.search(
        r"renegotiation.*?(after|within).*?\d+\s+(months|days)",
        text, re.IGNORECASE
    )
    renegotiation_terms = renegotiation_match.group(0) if renegotiation_match else None

    return jsonify({
        "text": text[:1000],  # Preview only
        "entities": entities,
        "tariffs": tariffs,
        "renegotiation_terms": renegotiation_terms,
        "pages": pages
    })

if __name__ == "__main__":
    app.run(port=5001, debug=True)
