const btn = document.getElementById('input-enviar')

function Login() {
    const data = new FormData()
    data.append('email', document.getElementById('input-email').value)
    data.append('senha', document.getElementById('input-senha').value)

    fetch("http://127.0.0.1:8080/api/login", {
        method: "POST",
        body: data
    }).then((e) => {

        e.json()
            .then((body) => {
                localStorage.setItem("jwt", body.JWT)
                window.location = 'sendmail.html'

            })

    })
}

btn.addEventListener('click', Login)