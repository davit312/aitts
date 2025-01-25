function getText(){
    return document.querySelector('#text').value
}
function play(file){
    let player = document.querySelector('#speech')
    let address = (atob(file)).split('\\')
    alert(address[address.length - 1])
    player.src = "/audio/"+address[address.length - 1]
    player.play()
}