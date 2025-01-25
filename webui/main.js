let audioQueue = []
let nextChunkToPlay = 0
let useCurrentQueue = false
let player = document.querySelector('#speech')

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
