var images = document.querySelectorAll(".image img")
var btnPrev = document.querySelector(".prev")
var btnNext = document.querySelector(".next")
var btnClose = document.querySelector(".close")
var gallery = document.querySelector(".gallery")
var galleryImg = document.querySelector(".gallery__inner img")

var currentImgIndex = 0;


function showGallery() {
    if (currentImgIndex == 0) {
        btnPrev.classList.add('hidden')
    }
    if (currentImgIndex == 7){
        btnNext.classList.add('hidden')
    } 
    galleryImg.src = images[currentImgIndex].src
    gallery.classList.add('show')
}

images.forEach((item, index) => {
    item.addEventListener("click", (evt) => {
        currentImgIndex = index;
        showGallery()
    })
})

btnClose.addEventListener('click', (evt) => {
    gallery.classList.remove("show")
})

document.addEventListener('keydown', (key) => {
    if (key.keyCode == 27) {
        gallery.classList.remove("show")
    }
})

btnNext.addEventListener('click', () => {
    if (currentImgIndex < images.length - 1) {
        currentImgIndex++
        showGallery()
    }
    btnPrev.classList.remove('hidden')
})

btnPrev.addEventListener('click', () => {
    if (currentImgIndex > 0) {
        currentImgIndex--
        showGallery()
    }
    btnNext.classList.remove('hidden')
})