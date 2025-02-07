function initSettings(){
    fetch('/webui/settings.json').then(resp => {
        resp.json().then(jr => {
            settings = jr

            // Run callbacks after settings loaded
            initModelSelector()
            if(settings.read_clipboard){
                readonClip.click()
            }
        })
    })
}

function initModelSelector() {
    let m = models.split(',')
    let optgroups = {}
 
    const UNKNOWN = "Unknown"
    let hasUnknown = false

    Array.from(m).forEach(el => {

        let model_lang = el.split("-")[0]
        
        let lang
        if(LANGUAGES[model_lang]){
            lang = LANGUAGES[model_lang]
        } else {
            lang = UNKNOWN
            hasUnknown = true
        }

        if(!optgroups[lang]){
            optgroups[lang] = []
        }

        let opt = document.createElement('option')
        opt.value = el
        opt.innerHTML = el

        if (el == settings.model){
            opt.selected = true
            setModel(settings.model)
        }

        optgroups[lang].push(opt)
    })

    let languages = Object.keys(optgroups)
    if (hasUnknown){
        languages = languages.filter( (ln) => ln !== UNKNOWN);
        languages.push(UNKNOWN)
    }

    languages.sort().forEach(l => {
        let group = document.createElement('optgroup')
        group.label = l

        optgroups[l].forEach(elem => {
            group.appendChild(elem)
        })
        modelSelector.appendChild(group)
    })
    
    modelSelector.addEventListener('change', (e) => {
        setModel(e.target.value)
        useCurrentQueue = false
    })
}

function init() {
    getModels() // Defined in go
}
