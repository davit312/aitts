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