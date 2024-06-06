const baseUrl = "http://localhost:8080"
//Elementos de card
const botonEliminar = document.getElementById("eliminar")
const totalTexto = document.getElementById("total")
const botonLogout = document.getElementById("logout")
const key = "cedula"

function createCard(auto) {
    return `
        <div class="card">
            <img class="car-image" src="${auto.img}" alt="${auto.marca}">
            <div class='car-details'>
            <p>Marca:<strong>${auto.marca}</strong> Sillas: <strong>${auto.sillas}</strong></p>
            <p>Combustible:<strong>${auto.combustible}</strong> Transmision: <strong>${auto.transmision}</strong></p>
            <p>Silla Bebe:<strong>${auto.sillabb}</strong> Seguros: <strong>${auto.seguros}</strong></p>
            <p>Lujos:<strong>${auto.lujos}</strong> Precio: <strong>${auto.precio}</strong></p>
            <button id='eliminar' onclick="eliminarReserva('${auto.id}')">Eliminar Reserva</button>
            </div>
        </div>
    `;
}
// Función para obtener datos del API y mostrar las tarjetas
async function loadCards() {

    try {
        const url = `${baseUrl}/autos/${localStorage.getItem(key)}`
        const response = await fetch(url);
        const data = await response.json();

        const cardsContainer = document.querySelector('.reservations');
        data.forEach(auto => {
            const cardHTML = createCard(auto);
            cardsContainer.innerHTML += cardHTML;
        });

        const total = data.reduce((total, producto) => total + producto.precio, 0)
        totalTexto.textContent = `${total}`
    } catch (error) {
        console.error('Error al obtener los datos:', error);
        if (localStorage.getItem("cedula") == null) {
            window.location.replace('/index.html')
        }

    }
}

window.onload = loadCards;

async function eliminarReserva(id) {
    const url = `${baseUrl}/reserva/${id}/0`;
    console.log(url)
    try {
        await fetch(url);
        window.location.reload();
    } catch (error) {
        console.log(error)
    }
}


botonLogout.addEventListener("click", async function (event) {
    event.preventDefault();
    localStorage.setItem("cedula", null)
});