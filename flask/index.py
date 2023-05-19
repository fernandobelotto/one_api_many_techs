from flask import Flask, request, jsonify
import random
import string

app = Flask(__name__)

notes = []

# Create a new note
@app.route('/notes', methods=['POST'])
def create_note():
    data = request.get_json()
    title = data.get('title')
    description = data.get('description')
    note_id = generate_id()
    new_note = {'id': note_id, 'title': title, 'description': description}
    notes.append(new_note)
    return jsonify(new_note), 201

# Read all notes
@app.route('/notes', methods=['GET'])
def get_notes():
    return jsonify(notes)

# Read a specific note
@app.route('/notes/<note_id>', methods=['GET'])
def get_note(note_id):
    note = next((note for note in notes if note['id'] == note_id), None)
    if note:
        return jsonify(note)
    return jsonify({'error': 'Note not found'}), 404

# Update a note
@app.route('/notes/<note_id>', methods=['PUT'])
def update_note(note_id):
    note = next((note for note in notes if note['id'] == note_id), None)
    if note:
        data = request.get_json()
        note['title'] = data.get('title')
        note['description'] = data.get('description')
        return jsonify(note)
    return jsonify({'error': 'Note not found'}), 404

# Delete a note
@app.route('/notes/<note_id>', methods=['DELETE'])
def delete_note(note_id):
    for note in notes:
        if note['id'] == note_id:
            notes.remove(note)
            return jsonify(note)
    return jsonify({'error': 'Note not found'}), 404

# Generate a random ID
def generate_id():
    id_length = 9
    characters = string.ascii_letters + string.digits
    return ''.join(random.choice(characters) for _ in range(id_length))

if __name__ == '__main__':
    app.run(debug=True)
