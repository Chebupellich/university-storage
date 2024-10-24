const openDialogButton = document.querySelectorAll('.add-task-icon');
const taskDialog = document.getElementById('taskDialog');
const closeDialogButton = document.getElementById('closeDialog');

const addTaskButton = document.getElementById('add-task-btn')
const text = document.getElementById('taskName')

let currCol = 0

openDialogButton.forEach(function (plus) {
    plus.addEventListener('click', () => {

        const input = plus.previousElementSibling

        if (input.value !== '') {
            const col = input.parentNode.parentNode.cellIndex
            const cell = document.getElementById('main-table').rows[2].cells[col]
            if (cell) {
                createNoteElement(input.value, cell.firstElementChild)
                input.value = ''
            }
            return
        }

        currCol = plus.parentNode.parentNode.cellIndex
        taskDialog.showModal();
        document.body.style.overflow = 'hidden';
        document.body.classList.add('modal-open');
    })
})

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
    const cell = document.getElementById('main-table').rows[2].cells[currCol]
    console.log(currCol, cell)
    if (cell) {
        createNoteElement(text.value, cell.firstElementChild)
    }
})