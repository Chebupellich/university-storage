const openDialogButton = document.getElementById('openDialog');
const taskDialog = document.getElementById('taskDialog');
const closeDialogButton = document.getElementById('closeDialog');

const addTaskButton = document.getElementById('add-task-btn')
const text = document.getElementById('taskName')

openDialogButton.addEventListener('click', () => {
    taskDialog.showModal();
    document.body.style.overflow = 'hidden';
    document.body.classList.add('modal-open');
});

closeDialogButton.addEventListener('click', () => {
    taskDialog.close();
    document.body.style.overflow = 'auto';
    document.body.classList.remove('modal-open');
});

taskDialog.addEventListener('click', (event) => {
    if (event.target === taskDialog) {
        taskDialog.close();
        document.body.style.overflow = 'auto';
        document.body.classList.remove('modal-open');
    }
});

addTaskButton.addEventListener('click', e => {
    const cell = document.getElementById('main-table').rows[2].cells[0]
    if (cell) {
        createNoteElement(text.value, cell.firstElementChild)
    }
})