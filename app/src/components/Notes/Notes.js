import React, { useEffect, useState } from 'react';
import goServer from '../../api/go-server';

function displayNotes(notes, selectNoteForEdit, deleteNote) {
  return (
    <ul className="list-group">
      {notes.map((note, index) => (
        <li className="list-group-item" key={index}>
          <h4>{note.title}</h4>
          <p style={{ wordBreak: 'break-all' }}>{note.content}</p>
          <button className="btn btn-info me-2" onClick={() => selectNoteForEdit(index)}>
            Modify
          </button>
          <button className="btn btn-danger" onClick={() => deleteNote(index)}>
            Delete
          </button>
        </li>
      ))}
    </ul>
  )
}

function notesForm(newNoteTitle, newNoteContent, setNewNoteTitle, setNewNoteContent, addNote, selectedNote) {
  return (
    <form>
      <div className="mb-3">
        <label className="form-label">Title:</label>
        <input type="text" className="form-control" value={newNoteTitle} onChange={(e) => setNewNoteTitle(e.target.value)} />
      </div>
      <div className="mb-3">
        <label className="form-label">Content:</label>
        <textarea className="form-control" value={newNoteContent} onChange={(e) => setNewNoteContent(e.target.value)} style={{ height: '150px' }} />
      </div>
      <button type="button" className={selectedNote !== null ? 'btn btn-info' : 'btn btn-success'} onClick={addNote} >
        {selectedNote !== null ? 'Update note' : 'Add a note'}
      </button>
    </form>
  )
}

function Notes() {
  const [notes, setNotes] = useState([]);
  const [newNoteTitle, setNewNoteTitle] = useState('');
  const [newNoteContent, setNewNoteContent] = useState('');
  const [showForm, setShowForm] = useState(false);
  const [selectedNote, setSelectedNote] = useState(null);

  const addNote = () => {
    if (newNoteTitle.trim() !== '' && newNoteContent.trim() !== '') {
      const newNote = {
        title: newNoteTitle,
        content: newNoteContent,
      };

      if (selectedNote !== null) {
        goServer.updateNoteForUser(parseInt(localStorage.getItem('current_user_id') || 0), notes[selectedNote].id, newNoteTitle, newNoteContent).then((response) => {
          newNote.id = response.data.note_id;
          const updatedNotes = [...notes];
          updatedNotes[selectedNote] = newNote;
          setNotes(updatedNotes);
          setSelectedNote(null);
          console.log(response.data.message);
        }).catch((error) => {
          console.log(error);
        });
      } else {
        goServer.createNote(localStorage.getItem('current_user_id'), newNoteTitle, newNoteContent).then((response) => {
          newNote.id = response.data.note_id;
          setNotes([...notes, newNote]);
          console.log(response.data.message);
        }).catch((error) => {
          console.log(error);
        });
      }

      setNewNoteTitle('');
      setNewNoteContent('');
      setShowForm(false);
    }
  };

  const selectNoteForEdit = (index) => {
    setSelectedNote(index);
    setNewNoteTitle(notes[index].title);
    setNewNoteContent(notes[index].content);
    setShowForm(true);
  };

  const deleteNote = (index) => {
    const noteID = notes[index].id;

    const updatedNotes = [...notes];
    updatedNotes.splice(index, 1);
    setNotes(updatedNotes);

    goServer.deleteNoteForUser(parseInt(localStorage.getItem('current_user_id') || 0), noteID).then((response) => {
      console.log(response.data.message);
    }).catch((error) => {
      console.log(error);
    });

    if (selectedNote === index) {
      setSelectedNote(null);
      setNewNoteTitle('');
      setNewNoteContent('');
      setShowForm(false);
    }
  };

  useEffect(() => {
    goServer.getNotesForUser(localStorage.getItem('current_user_id')).then((response) => {
      const mappedNotes = response.data.map((note) => {
        return {
          id: note.ID,
          title: note.Title,
          content: note.Content,
        };
      });

      setNotes(mappedNotes);
    }).catch((error) => {
      console.log(error);
    });
  }, []);

  return (
    <div className="container mt-4">
      <h2 className="mb-4">Notes</h2>

      <button className="btn btn-primary mb-3" onClick={() => { setShowForm(!showForm); setSelectedNote(null); setNewNoteTitle(''); setNewNoteContent(''); }}>
        {showForm ? 'Hide form' : 'Add a note'}
      </button>

      {showForm && notesForm(newNoteTitle, newNoteContent, setNewNoteTitle, setNewNoteContent, addNote, selectedNote)}

      {!showForm && displayNotes(notes, selectNoteForEdit, deleteNote)}
    </div>
  );
}

export default Notes;
