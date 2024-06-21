document.addEventListener('DOMContentLoaded', () => {
    const noteForm = document.getElementById('noteForm');
    const noteInput = document.getElementById('noteInput');
    const notesList = document.getElementById('notesList');

    noteForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const content = noteInput.value;
        const response = await fetch('/notes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ content })
        });
        if (response.ok) {
            noteInput.value = '';
            loadNotes();
        }
    });

    async function loadNotes() {
        const response = await fetch('/notes');
        const notes = await response.json();
        notesList.innerHTML = notes.map(note => `<li>${note}</li>`).join('');
    }

    loadNotes();
});
