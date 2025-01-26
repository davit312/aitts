function initModels(){
    fetch('/webui/settings.json').then(resp => {
        resp.json().then(jr => {
            settings = jr

            // Run callbacks after settings loaded
            initModelSelector()
        })
    })
}

function initModelSelector() {
    let m = models.split(',')
    Array.from(m).forEach(el => {
        let opt = document.createElement('option')
        opt.value = el
        opt.innerHTML = el
        if (el == settings.model){
            opt.selected = true
            setModel(settings.model)
        }
        modelSelector.appendChild(opt)
    })

    modelSelector.addEventListener('change', (e) => {
        setModel(e.target.value)
    })
}

function init() {
    getModels() // Defined in go
}
