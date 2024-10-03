const table = document.getElementById("main-table")

for (let i = 0; i < localStorage.length; i++) {
    const key = window.localStorage.key(i);
    const value = JSON.parse(window.localStorage.getItem(key));
    console.log("CREATE?: ", i, key, value)
    createNote(key, value.title, value.content, value.state, table.rows[2].cells[value.column].firstElementChild)
}

function createNote(key, title, content, state, noteList) {
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

    const keyInfo = document.createElement('div');
    keyInfo.className = 'key-block';
    keyInfo.textContent = key

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
}