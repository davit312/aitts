function initSettings(){
    fetch('/webui/settings.json').then(resp => {
        resp.json().then(jr => {
            settings = jr

            // Run callbacks after settings loaded
            initModelSelector()
            if(settings.read_clipboard){
                readonClip.click()
            }
            if(!models){
                if(confirm("No models found. Do you want to open setting and download some model?")){
                    document.getElementById("settings").click()
                }
            }
        })
    })
}

function initModelSelector() {
    let m = models.split(',')
    let optgroups = {}
 
    const UNKNOWN = "Unknown"
    let hasUnknown = false
    let model_found = false

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

        if (el == settings.default_model){
            opt.selected = true
            model_found = true
            setModel(settings.default_model)
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

    if(!model_found && m.length > 0){
        setModel(modelSelector.value)
    }
}

function init() {
    getModels() // Defined in go
}
