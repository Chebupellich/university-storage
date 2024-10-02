import { editElement, toggleComplete } from "./edit-notes"

function setupNoteListeners(link) {
    const titleElement = link.querySelector("h2");

    link.addEventListener("click", function (e) {
        if (e.target.tagName !== "IMG" && e.target.tagName !== "INPUT") {
            e.preventDefault()
        } else if (e.target.id === "card-checkbox") {
            toggleComplete(e.target)
        }
    })

    titleElement.addEventListener("click", function () {
        editElement(titleElement);
    });

    const contentElement = link.querySelector("p");
    contentElement.addEventListener("click", function () {
        editElement(contentElement);
    });
}

const createField = document.querySelectorAll(".create-task input")
createField.forEach(function (field) {
    field.addEventListener('keypress', function (e) {
        if (e.key === 'Enter') {
            const col = field.parentNode.parentNode.cellIndex

            const cell = document.getElementById('main-table').rows[2].cells[col]
            if (cell) {
                createNoteElement(field.value, cell.firstElementChild)
            }
        }
    })
})

function createNoteElement(title, noteList) {
    const li = document.createElement('li');

    const a = document.createElement('a');
    a.href = '#';
    a.className = 'note-link';

    const h2 = document.createElement('h2');
    h2.textContent = title;

    const p = document.createElement('p');

    const label = document.createElement('label');
    label.className = 'customCheckbox';

    const input = document.createElement('input');
    input.type = 'checkbox';
    input.id = 'card-checkbox';
    input.checked = 'false'

    const span = document.createElement('span');
    span.className = 'checkboxImage';

    const uncheckedImg = document.createElement('img');
    uncheckedImg.src = 'assets/icons/circle-outline.svg';
    uncheckedImg.alt = 'Incomplete';
    uncheckedImg.className = 'uncheckedImage';

    const checkedImg = document.createElement('img');
    checkedImg.src = 'assets/icons/check-circle-outline.svg';
    checkedImg.alt = 'Complete';
    checkedImg.className = 'checkedImage';

    let uniqueKey;
    do {
        uniqueKey = `item_${Math.random().toString(36).substring(2, 9)}`;
    } while (localStorage.getItem(uniqueKey) !== null);

    const keyInfo = document.createElement('div');
    keyInfo.className = 'key-block';
    keyInfo.textContent = uniqueKey

    span.appendChild(uncheckedImg);
    span.appendChild(checkedImg);
    label.appendChild(input);
    label.appendChild(span);
    a.appendChild(h2);
    a.appendChild(p);
    a.appendChild(label);
    a.appendChild(keyInfo)
    li.appendChild(a);

    noteList.appendChild(li);

    const data = {
        title: title || '#' + itemKey,
        content: '',
        column: noteList.parentNode.cellIndex,
        state: "false"
    };

    window.localStorage.setItem(uniqueKey, JSON.stringify(data));
    setupNoteListeners(a)
}