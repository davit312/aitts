Array.from(document.querySelectorAll('.voice-btn')).forEach(el => {
    let langElem = el.parentElement.parentElement.parentElement.parentElement.parentElement

    let code = langElem.querySelector('code').innerText
    let quality = el.parentElement.innerText.split(' - ')[0]
    let voice = el.parentElement.parentElement.parentElement.innerText.split('\n')[0]

    let modelID = `${code}-${voice}-${quality}`

    el.addEventListener('click', e => {
        let links = []
        Array.from(e.target.parentElement.querySelectorAll('span')).forEach(el => {
        links.push(el.dataset.href)
        })
        alert(links.join('\n'))
        alert(modelID)
    })
})
