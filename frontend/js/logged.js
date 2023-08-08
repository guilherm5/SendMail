const token = localStorage.getItem('jwt')
const logged = document.getElementById("logged")

fetch("http://127.0.0.1:8080/api/logged", {
    method: "GET",
    headers: {
        'Authorization': localStorage.getItem('jwt')
    }
}).then((e) => {
    if (e.status === 401) {
        window.location = 'login.html'
    }
    return e.json()
}).then((body) => {
    const obj = JSON.stringify(body)
    console.log(obj)
    logged.innerHTML = obj
})
