const express = require('express');
const app = express();
app.use(express.json());

const notes = [];

// Create a new note
app.post('/notes', (req, res) => {
  const { title, description } = req.body;
  const id = generateId();
  const newNote = { id, title, description };
  notes.push(newNote);
  res.status(201).json(newNote);
});

// Read all notes
app.get('/notes', (req, res) => {
  res.json(notes);
});

// Read a specific note
app.get('/notes/:id', (req, res) => {
  const { id } = req.params;
  const note = notes.find((note) => note.id === id);
  if (!note) {
    res.status(404).json({ error: 'Note not found' });
  } else {
    res.json(note);
  }
});

// Update a note
app.put('/notes/:id', (req, res) => {
  const { id } = req.params;
  const { title, description } = req.body;
  const note = notes.find((note) => note.id === id);
  if (!note) {
    res.status(404).json({ error: 'Note not found' });
  } else {
    note.title = title;
    note.description = description;
    res.json(note);
  }
});

// Delete a note
app.delete('/notes/:id', (req, res) => {
  const { id } = req.params;
  const index = notes.findIndex((note) => note.id === id);
  if (index === -1) {
    res.status(404).json({ error: 'Note not found' });
  } else {
    const deletedNote = notes.splice(index, 1);
    res.json(deletedNote[0]);
  }
});

// Generate a random ID
function generateId() {
  return Math.random().toString(36).substr(2, 9);
}

// Start the server
app.listen(3000, () => {
  console.log('Server is running on port 3000');
});
