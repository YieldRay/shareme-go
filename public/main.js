import { ShareMe, isNamespaceValid } from "./client.js";
const app = new ShareMe(location.origin);
window.app = app;
console.log(
    "%c     _____ _     %c               __  __      \n" +
        "%c    / ____| |%c                  |  \\/  |     \n" +
        "%c   | (___ | |%c__   __ _ _ __ ___| \\  / | ___ \n" +
        "%c    \\___ \\|%c '_ \\ / _` | '__/ _ \\ |/\\| |/ _ \\ \n" +
        "%c    ____) | |%c | | (_| | | |  __/ |  | |  __/\n" +
        "%c   |_____/|_|%c |_|\\__,_|_|  \\___|_|  |_|\\___|\n" +
        `%c                           v${ShareMe.VERSION}-by.YieldRay`,

    "color:#ff0000",
    "color:#ff0000",
    "color:#ff3b00",
    "color:#ff7500",
    "color:#ff7800",
    "color:#FD7B00",
    "color:#FFAD00",
    "color:#FEDA00",
    "color:#D0FD00",
    "color:#93FF00",
    "color:#80FF00",
    "color:#1AFF00",
    "color:#00FF2E"
);
console.log(
    `%c

Usage for command line  (replace \`:namespace\` with a namespace you want)  
        
$ curl ${location.origin}/:namespace                                              
$ curl ${location.origin}/:namespace -d t=any_thing_you_want_to_store

`,
    "color: #66ccff; font-size: 16px; padding: 2px;"
);

// forbid invalid namespace
const namespace = window.location.pathname.slice(1);
if (!isNamespaceValid(namespace)) location.pathname = generateRandomString();

function generateRandomString(length = 4) {
    const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let text = "";
    for (let i = 0; i < length; i++) text += possible.charAt(Math.floor(Math.random() * possible.length));
    return text;
}

async function updateFromServer($textarea, $info) {
    $info.className = "blue";
    $textarea.disabled = true;
    const data = await app.get(namespace);
    $textarea.disabled = false;
    if (data === null) {
        $info.className = "red";
    } else {
        $info.className = "green";
        $textarea.value = data;
        $textarea.focus();
    }
    $info.innerText = "Updated at: " + new Date().toLocaleString();
}

async function updateToServer($textarea, $info) {
    $info.className = "blue";
    const data = $textarea.value;
    const success = await app.set(namespace, data);
    $info.className = success ? "green" : "red";
    $info.innerText = "Updated at: " + new Date().toLocaleString();
}

function debounce(func, wait = 0) {
    let timer = null;
    return function (...args) {
        clearTimeout(timer);
        timer = setTimeout(() => func(...args), wait);
    };
}

document.addEventListener("DOMContentLoaded", () => {
    const $textarea = document.getElementById("$textarea");
    const $info = document.getElementById("$info");
    updateFromServer($textarea, $info);
    $info.addEventListener("click", () => updateFromServer($textarea, $info));
    $textarea.addEventListener(
        "input",
        debounce(() => updateToServer($textarea, $info), 1000)
    );
});
