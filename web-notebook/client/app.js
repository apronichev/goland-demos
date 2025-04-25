document.addEventListener('DOMContentLoaded', () => {
    const noteForm = document.getElementById('noteForm');
    const noteInput = document.getElementById('noteInput');
    const notesList = document.getElementById('notesList');

    noteForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        const content = noteInput.value.trim();

        if (!content) return; // Don't submit empty notes

        const response = await fetch('/notes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ content })
        });

        if (response.ok) {
            noteInput.value = '';
            await loadNotes();
        } else {
            console.error('Failed to add note');
        }
    });

    async function loadNotes() {
        try {
            const response = await fetch('/notes');
            const notes = await response.json();

            if (Array.isArray(notes)) {
                notesList.innerHTML = notes
                    .map((note, index) => `<li>Note ${index + 1}: ${note.content}</li>`)
                    .join('');
            } else {
                console.error('Unexpected response format:', notes);
            }
        } catch (error) {
            console.error('Error loading notes:', error);
            notesList.innerHTML = '<li>Error loading notes</li>';
        }
    }

    loadNotes();
});