function toggleBurgerMenu() {
    var x = document.getElementById("header-nav");
    if (x.className === "main-nav") {
        x.className += " responsive";
    } else {
        x.className = "main-nav";
    }
}