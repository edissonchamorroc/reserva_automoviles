
const baseUrl = "http://localhost:8080"
//Elementos login
const cedula = document.querySelector("#cedula");
const contrasena = document.querySelector("#contrasena")
const botonLogin = document.querySelector("#login")


botonLogin.addEventListener("click", async function (event) {
    event.preventDefault();

    const url = `${baseUrl}/login/${cedula.value}/${contrasena.value}`
    const response = await fetch(url);
    const data = await response.json();
    if (data.cedula != null) {
        localStorage.setItem("cedula", data.cedula);
        console.log(localStorage.getItem("cedula"))
        window.location.replace('/home.html')
    } else {
        window.location.replace('/index.html')
    }

});



