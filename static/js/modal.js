var btnOpen = document.querySelector(".open-modal-btn")
var modal = document.querySelector(".modal")
var btnClose = document.querySelector(".modal__footer button")
var iconClose = document.querySelector(".modal__header i")





btnOpen.addEventListener("click", (evt)=>{
    modal.classList.toggle("hide")
})

modal.addEventListener("click",(evt)=>{
    if (evt.target != evt.currentTarget){
        return
    }
    modal.classList.toggle("hide")
})
iconClose.addEventListener("click", (evt)=>{
    modal.classList.toggle("hide")
})
btnClose.addEventListener("click", (evt)=>{
    modal.classList.toggle("hide")
})



