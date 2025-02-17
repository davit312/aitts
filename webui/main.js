const LANGUAGES = {
    "hy_AM": "Armenian",
    "en_US": "English (US)",
    "ru_RU": "Russian"
}

let settings = {}
let models = []
let currentModel = ""
let audioQueue = []
let nextChunkToPlay = 0
let useCurrentQueue = false
let userPaused = false
let player = document.querySelector('#speech')
let textbox = document.querySelector('#text')
let modelSelector = document.querySelector('#models')
let readonClip = document.querySelector('#readonclipboard')

readonClip.addEventListener('change', (e) => {
    setClipTrack(e.target.checked)
})

textbox.addEventListener('change', (e) => {
    useCurrentQueue = false
})

player.addEventListener('ended', (e) => {
    if (nextChunkToPlay > audioQueue.length -1){
        return
    }
    play(audioQueue[nextChunkToPlay])
    nextChunkToPlay += 1
})

// Get the modal
var modal = document.getElementById("settingsModal");

// Get the button that opens the modal
var btn = document.getElementById("settings");

// Get the <span> element that closes the modal
var span = document.getElementsByClassName("close")[0];

// When the user clicks the button, open the modal
btn.onclick = function() {

    fetch('/webui/languages.html')
        .then(response => response.text())
        .then(data => {
            let langDiv = modal.querySelector('#langlist')
            langDiv.innerHTML = data;
            
            let script = document.createElement('script');
            script.src = '/webui/modalLangs.js';
            langDiv.appendChild(script);
        })

    modal.style.display = "block";

    let readonstart = document.querySelector('#readonstart')
    let defmodel = document.querySelector('#defaultmodel')
    
    defmodel.innerHTML = document.querySelector('#models').innerHTML
    defmodel.value = settings.default_model

    readonstart.checked = false
    if(settings.read_clipboard){
        readonstart.click()
    }
}

// When the user clicks on <span> (x), close the modal
span.onclick = function() {
    modal.style.display = "none";
}

// When the user clicks anywhere outside of the modal, close it
window.onclick = function(event) {
    if (event.target == modal) {
        modal.style.display = "none";
    }
}

// Save settings
document.querySelector('#saveDefaultSettings').addEventListener('click', (e) => {
    settings.default_model = document.querySelector('#defaultmodel').value
    settings.read_clipboard = document.querySelector('#readonstart').checked
    saveSettings(JSON.stringify(settings))
    alert('Settings saved')
})

function getText(){
    return document.querySelector('#text').value
}

function read(){
    if(userPaused){
        userPaused = false
        player.play()
    } else {
        if (useCurrentQueue){
            nextChunkToPlay = 1
            play(audioQueue[0])
        } else{
            readText(getText())
        }
    }
}

function play(chunk){
    if (userPaused){
        return
    }
    player.src = "/audio/"+chunk
    player.play()
}

function addToQueue(chunk, start=false){
    if(start){
        audioQueue = []
        nextChunkToPlay = 0
        player.pause()
        useCurrentQueue = true
    }
    audioQueue.push(chunk)
    if (start || player.paused){
        if(!userPaused){
            play(audioQueue[nextChunkToPlay])
            nextChunkToPlay += 1
        }
    }
}

function pause(){
    if (player.paused){
        return
    }
    userPaused = true
    player.pause()
}

function stop(){
    player.pause()
    nextChunkToPlay = 0
    useCurrentQueue = false
    userPaused = false
}

/* Start program UI */
init()