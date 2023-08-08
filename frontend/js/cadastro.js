const btn = document.getElementById('input-enviar')

function CadastroUsuario() {
    const data = new FormData()
    data.append('nome', document.getElementById('input-nome').value)
    data.append('email', document.getElementById('input-email').value)
    data.append('senha', document.getElementById('input-senha').value)

    fetch("http://127.0.0.1:8080/user", {
        method: 'POST',
        body: data
    })
        .then((e) => {
            console.log(e.status)
        })
        .then((body) => {
            console.log(body)
            window.location = '/frontend/html/login.html'

        })
}

btn.addEventListener('click', CadastroUsuario)

