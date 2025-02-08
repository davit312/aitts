let countdown = 0

function modelActionSuccess() {
    countdown--
    if (countdown == 0) {
        location.reload()
    }
}

function modelActionFailure(message = '') {
    countdown = 0
    alert('Error: ' + message)
    location.reload()
}

Array.from(document.querySelectorAll('.voice-btn')).forEach(el => {
    let langElem = el.parentElement.parentElement.parentElement.parentElement.parentElement

    let code = langElem.querySelector('code').innerText
    let quality = el.parentElement.innerText.split(' - ')[0]
    let voice = el.parentElement.parentElement.parentElement.innerText.split('\n')[0]

    let modelID = `${code}-${voice}-${quality}`
    let action = ''

    if (models.indexOf(modelID) !== -1) {
        el.innerText = 'Remove'
        action = 'remove'
    } else {
        el.innerText = 'Download'
        action = 'download'
    }

    el.addEventListener('click', e => {
        if (action == 'download'){
            let links = []
            Array.from(e.target.parentElement.querySelectorAll('span')).forEach(el => {
                links.push(el.dataset.href)
            })

            countdown = links.length
            links.forEach(link => onModelAction(action, link))
        } else if (action == 'remove') {
            if(!confirm("Do you want to delete model: \n" + modelID)){
                return
            }
            countdown = 2
            onModelAction(action, modelID+'.onnx')
            onModelAction(action, modelID+'.onnx.json')
        }
    })
})
