document.addEventListener("DOMContentLoaded", function () {
    const notes = document.querySelectorAll("li a");

    notes.forEach(function (link) {
        setupNoteListeners(link)
    });
});

function editElement(element) {
    const currentText = element?.textContent;
    const textarea = document.createElement("textarea");

    textarea.value = currentText || '';
    element.replaceWith(textarea);
    textarea.focus();

    textarea.addEventListener("blur", function () {
        saveEdit(textarea, element);
    });

    textarea.addEventListener("keypress", function (e) {
        if (e.key === "Enter" && !e.shiftKey) {
            saveEdit(textarea, element);
        } else if (e.key === "Enter" && !e.shiftKey) {
            textarea.value += "\n"
        }
    });
}

function toggleComplete(input) {
    const itemKey = input.parentNode.parentNode.parentNode.lastElementChild.textContent
    const data = JSON.parse(localStorage.getItem(itemKey))

    localStorage.setItem(itemKey, JSON.stringify(data))

    if (data.state === 'false') {
        input.setAttribute('checked', true)
        data.state = 'true'
    } else {
        input.setAttribute('checked', false)
        data.state = 'false'
    }
    localStorage.setItem(itemKey, JSON.stringify(data))
}

function saveEdit(textarea, originalElement) {
    const newText = textarea.value;
    const newElement = document.createElement(originalElement.tagName);
    newElement.textContent = newText;

    textarea.replaceWith(newElement);
    newElement.addEventListener("click", function () {
        editElement(newElement);
    });

    if (newElement.parentNode) {
        const link = newElement.parentNode.parentNode
        const noteTitle = link.querySelector("h2").textContent;
        const noteContent = link.querySelector("p").textContent;

        const itemKey = newElement.parentNode.lastElementChild.textContent

        const data = {
            title: noteTitle || '',
            content: noteContent || '',
            column: link.parentNode.parentNode.cellIndex,
            state: "false"
        };

        window.localStorage.setItem(itemKey, JSON.stringify(data));
    }
}

const createField = document.querySelectorAll(".create-task input")
createField.forEach(function (field) {
    field.addEventListener('keypress', function (e) {
        if (e.key === 'Enter' && field.value) {
            const col = field.parentNode.parentNode.cellIndex

            const cell = document.getElementById('main-table').rows[2].cells[col]
            if (cell) {
                createNoteElement(field.value, cell.firstElementChild)
                field.value = ''
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

    const groupImg = document.createElement('div')
    groupImg.className = 'noteImageGroup'

    const uncheckedImg = document.createElement('img');
    uncheckedImg.src = 'assets/icons/circle-outline.svg';
    uncheckedImg.alt = 'Incomplete';
    uncheckedImg.className = 'uncheckedImage';

    const checkedImg = document.createElement('img');
    checkedImg.src = 'assets/icons/check-circle-outline.svg';
    checkedImg.alt = 'Complete';
    checkedImg.className = 'checkedImage';

    const thrashCan = document.createElement('img');
    thrashCan.src = 'assets/icons/delete.svg'
    thrashCan.alt = 'deleteNote'
    thrashCan.className = 'deleteImage'

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

    groupImg.appendChild(label)
    groupImg.appendChild(thrashCan)

    a.appendChild(h2);
    a.appendChild(p);
    a.appendChild(groupImg);
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

function setupNoteListeners(link) {
    const titleElement = link.querySelector("h2");

    link.addEventListener("click", function (e) {
        if (e.target.tagName !== "IMG" && e.target.tagName !== "INPUT") {
            e.preventDefault()
        } else if (e.target.id === "card-checkbox") {
            toggleComplete(e.target)
        } else if (e.target.className === "deleteImage") {
            deleteNote(e.target)
        }
    })

    titleElement.addEventListener("click", function () {
        editElement(titleElement);
    });

    const contentElement = link.querySelector("p");
    contentElement.addEventListener("click", function () {
        editElement(contentElement);
    });

    link.addEventListener("dragstart", (e) => {
        e.dataTransfer.setData("text/plain", e.target.lastElementChild.textContent);
    });

    link.addEventListener('dragend', (e) => {
        const itemKey = e.target.lastElementChild.textContent;
        if (localStorage.getItem(itemKey) === null) {
            link.parentNode.remove();
            link.remove()
        }
        e.preventDefault()
    })
}

function deleteNote(link) {
    const key = link.parentNode.parentNode.lastElementChild.textContent
    localStorage.removeItem(key)

    link.parentNode.parentNode.parentNode.remove()
}

const columns = document.querySelectorAll(".note-row td")
columns.forEach(function (col) {
    col.addEventListener('dragover', function (e) {
        e.preventDefault()
    })

    col.addEventListener('drop', function (e) {
        const key = e.dataTransfer.getData("text/plain")
        const data = JSON.parse(localStorage.getItem(key))

        if (data && col.contains(e.target)) {
            createDragNote(data.title, data.content, data.state, key, col.firstElementChild)
        }
    })
})

function createDragNote(title, content, state, key, noteList) {
    const li = document.createElement('li');
    const a = document.createElement('a');
    a.href = '#';
    a.className = 'note-link';

    const h2 = document.createElement('h2');
    h2.textContent = title;

    const p = document.createElement('p');
    p.textContent = content

    const label = document.createElement('label');
    label.className = 'customCheckbox';

    const input = document.createElement('input');
    input.type = 'checkbox';
    input.id = 'card-checkbox';
    input.checked = state === 'true' ? false : true

    const span = document.createElement('span');
    span.className = 'checkboxImage';

    const groupImg = document.createElement('div')
    groupImg.className = 'noteImageGroup'

    const uncheckedImg = document.createElement('img');
    uncheckedImg.src = 'assets/icons/circle-outline.svg';
    uncheckedImg.alt = 'Incomplete';
    uncheckedImg.className = 'uncheckedImage';

    const checkedImg = document.createElement('img');
    checkedImg.src = 'assets/icons/check-circle-outline.svg';
    checkedImg.alt = 'Complete';
    checkedImg.className = 'checkedImage';

    const thrashCan = document.createElement('img');
    thrashCan.src = 'assets/icons/delete.svg'
    thrashCan.alt = 'deleteNote'
    thrashCan.className = 'deleteImage'

    window.localStorage.removeItem(key)
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

    groupImg.appendChild(label)
    groupImg.appendChild(thrashCan)

    a.appendChild(h2);
    a.appendChild(p);
    a.appendChild(groupImg);
    a.appendChild(keyInfo)
    li.appendChild(a);

    noteList.appendChild(li);

    const data = {
        title: title || '#' + key,
        content: content || '',
        column: noteList.parentNode.cellIndex,
        state: state
    };
    window.localStorage.setItem(uniqueKey, JSON.stringify(data));
    setupNoteListeners(a)
}