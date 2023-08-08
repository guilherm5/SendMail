const btn = document.getElementById("input-enviar");

async function sendMail() {
    const data = new FormData();
    data.append("from", document.getElementById("input-from").value);
    data.append("to", document.getElementById("input-to").value);
    data.append("assunto", document.getElementById("input-assunto").value);
    data.append("message", document.getElementById("message").value);

    try {
        const response = await fetch("http://127.0.0.1:8080/api/mail", {
            method: "POST",
            body: data,
            headers: {
                'Authorization': localStorage.getItem('jwt')
            }
        });

        if (response.status !== 200) {
            window.location = 'login.html';
            return;
        }

        console.log(response.status);
        window.location = 'sendMail.html';
    } catch(error) {
        console.error("Error occurred:", error);

    }
}

btn.addEventListener('click', sendMail);
