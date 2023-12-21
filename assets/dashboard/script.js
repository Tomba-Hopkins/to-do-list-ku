card = document.querySelector('.card')

card.addEventListener('click', function(){
    card.classList.toggle('flipcard')
})

todo1 = document.querySelector('.todo1')
todo2 = document.querySelector('.todo2')
melayang1 = document.querySelector('.melayang1')
melayang2 = document.querySelector('.melayang2')
close1 = document.querySelector('.close1')
close2 = document.querySelector('.close2')

todo1.addEventListener('click', function(){
    melayang1.style.display = 'flex'
})
todo2.addEventListener('click', function(){
    melayang2.style.display = 'flex'
})

close1.addEventListener('click', function(){
    melayang1.style.display = 'none'
})
close2.addEventListener('click', function(){
    melayang2.style.display = 'none'
})


// Input 

plus1 = document.querySelector(".plus1")
plus2 = document.querySelector(".plus2")
hiddenInput1 = document.querySelector(".hidden-input1")
hiddenInput2 = document.querySelector(".hidden-input2")
kirim1 = document.querySelector(".kirim1")
kirim2 = document.querySelector(".kirim2")

plus1.addEventListener('click', function(){
    hiddenInput1.style.display = "inline"
    kirim1.style.display = "inline"
})

plus2.addEventListener('click', function(){
    hiddenInput2.style.display = "inline"
    kirim2.style.display = "inline"

})

kirim1.addEventListener('click', function(){
    hiddenInput1.style.display = "none"
    kirim1.style.display = "none"
})
kirim2.addEventListener('click', function(){
    hiddenInput2.style.display = "none"
    kirim2.style.display = "none"
})