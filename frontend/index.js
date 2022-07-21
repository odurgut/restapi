getButton = document.getElementById("button");
idBox = document.getElementById("id");
titleBox = document.getElementById("title");
contentBox = document.getElementById("content");
authorBox = document.getElementById("author");

eventListeners();

function eventListeners() {
    getButton.addEventListener("click", getData);
}

function getData() {
    fetch("http://localhost:9000/api/posts")
        .then((response) => response.json())
        .then((json) => {
            console.log(json);
            updateUI(json);
        });
}

function updateUI(jsonData) {
    idBox.textContent = jsonData[0].id;
    titleBox.textContent = jsonData[0].title;
    contentBox.textContent = jsonData[0].content;
    authorBox.textContent = jsonData[0].author;
}
