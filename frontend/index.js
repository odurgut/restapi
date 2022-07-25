getButton = document.getElementById("button");
idBox = document.getElementById("id");
titleBox = document.getElementById("title");
contentBox = document.getElementById("content");
authorBox = document.getElementById("author");

postButton = document.getElementById("buttonPost");
title2Box = document.getElementById("title2");
content2Box = document.getElementById("content2");
author2Box = document.getElementById("author2");

eventListeners();

function eventListeners() {
    getButton.addEventListener("click", getData);
    postButton.addEventListener("click", postData);
}

function getData() {
    fetch("http://localhost:9000/api/posts")
        .then((response) => response.json())
        .then((json) => {
            console.log(json);
            updateUI(json);
        });
}

function postData() {
    const post = {
        title: title2Box.value,
        content: content2Box.value,
        author: author2Box.value,
    };

    fetch("http://localhost:9000/api/posts", {
        method: "POST",
        body: JSON.stringify(post),
    })
        .then((response) => response.json())
        .then((json) => {
            console.log(json);
        });
}

function updateUI(jsonData) {
    idBox.textContent = jsonData[jsonData.length - 1].id;
    titleBox.textContent = jsonData[jsonData.length - 1].title;
    contentBox.textContent = jsonData[jsonData.length - 1].content;
    authorBox.textContent = jsonData[jsonData.length - 1].author;
}
