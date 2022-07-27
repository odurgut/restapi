getButton = document.getElementById("button");
idBox = document.getElementById("id");
titleBox = document.getElementById("title");
contentBox = document.getElementById("content");
authorBox = document.getElementById("author");

postButton = document.getElementById("buttonPost");
title2Box = document.getElementById("title2");
content2Box = document.getElementById("content2");
author2Box = document.getElementById("author2");

clearButton = document.getElementById("clearButton");

eventListeners();

function eventListeners() {
    getButton.addEventListener("click", getData);
    postButton.addEventListener("click", postData);
    clearButton.addEventListener("click", clearTable);
}

function getData() {
    fetch("http://localhost:9000/api/posts")
        .then((response) => response.json())
        .then((json) => {
            console.log(json);
            for (var i = 0; i < json.length; i++) {
                updateUI(json, i);
            }
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
            getData();
        });
}

function updateUI(jsonData, i) {
    var tablePosts = document.getElementById("tablePosts");

    idBox.textContent = jsonData[jsonData.length - 1].id;
    titleBox.textContent = jsonData[jsonData.length - 1].title;
    contentBox.textContent = jsonData[jsonData.length - 1].content;
    authorBox.textContent = jsonData[jsonData.length - 1].author;

    var rowPosts = tablePosts.insertRow(1);
    var td11 = rowPosts.insertCell(0);
    var td12 = rowPosts.insertCell(1);
    var td13 = rowPosts.insertCell(2);
    var td14 = rowPosts.insertCell(3);

    td11.innerHTML = jsonData[i].id;
    td12.innerHTML = jsonData[i].title;
    td13.innerHTML = jsonData[i].content;
    td14.innerHTML = jsonData[i].author;
}

function clearTable() {
    var tablePosts = document.getElementById("tablePosts");

    var rowCount = tablePosts.rows.length;
    for (var x = rowCount - 1; x > 0; x--) {
        tablePosts.deleteRow(x);
    }
}
