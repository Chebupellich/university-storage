const dropdownNav = document.getElementById("drop-nav");
const burgerIcon = document.getElementById("burger-menu");

function toggleBurgerMenu() {
    dropdownNav.classList.toggle("clicked");
    burgerIcon.classList.toggle("clicked");
}

window.addEventListener('resize', function () {
    if (window.innerWidth > 1024) {
        dropdownNav.classList.remove("clicked");
        burgerIcon.classList.remove("clicked");
    }
});

document.addEventListener('click', function (event) {
    if (!dropdownNav.contains(event.target) && !burgerIcon.contains(event.target)) {
        dropdownNav.classList.remove("clicked");
        burgerIcon.classList.remove("clicked");
    }
});


const index = window.location.pathname.lastIndexOf('/')
const curr = window.location.pathname.substring(index + 1)
const links = document.querySelectorAll(".main-nav a, .dropdown-nav a");

links.forEach(function (link) {
    let temp = link.href.substring(link.href.lastIndexOf('/') + 1)
    if (temp == curr) {
        link.style.color = '#fff'
        link.style.textDecoration = "white underline"
    }

    link.addEventListener('click', function (e) {
        const index = window.location.pathname.lastIndexOf('/')
        const curr = window.location.pathname.substring(index + 1)

        let temp = link.href.substring(link.href.lastIndexOf('/') + 1)
        if (temp == curr) {
            e.preventDefault()
        }
    })
})