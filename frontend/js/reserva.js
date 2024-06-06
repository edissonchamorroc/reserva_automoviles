const baseUrl = "http://localhost:8080"
const botonAgregar = document.getElementById("agregar")
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
            <button id='agregar' onclick="agregarReserva('${auto.id}')">Reservar</button>
            </div>
        </div>
    `;
}
// Función para obtener datos del API y mostrar las tarjetas
async function loadCards() {
 
    try {
        const url = `${baseUrl}/autos/0`
        const response = await fetch(url);
        const data = await response.json();

        const cardsContainer = document.querySelector('.reservations');
        data.forEach(auto => {
            const cardHTML = createCard(auto);
            cardsContainer.innerHTML += cardHTML;
        });

    } catch (error) {
        console.error('Error al obtener los datos:', error);
        window.location.replace('/index.html')
    }
}

window.onload = loadCards;

async function agregarReserva(id) {
    const url = `${baseUrl}/reserva/${id}/${localStorage.getItem(key)}`;
    console.log(url)
    try {
        await fetch(url);
        window.location.reload();
    } catch (error) {
        console.log(error)
    }

}