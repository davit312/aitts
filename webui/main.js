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
let player = document.querySelector('#speech')
let modelSelector = document.querySelector('#models')

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

function play(chunk){
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
        play(audioQueue[nextChunkToPlay])
        nextChunkToPlay += 1
    }
}

/* Start program UI */
init()