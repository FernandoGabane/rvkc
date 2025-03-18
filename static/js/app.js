
var btnSignin = document.querySelector("#signin");
var btnSignup = document.querySelector("#signup");

var body = document.querySelector("body");


btnSignin.addEventListener("click", function () {
   body.className = "sign-in-js"; 
});

btnSignup.addEventListener("click", function () {
    body.className = "sign-up-js";
})

async function fetchPilot() {
    try {
        const res = await fetch('http:/localhost:8080/pilots/123456789');
        if (!res.ok) throw new Error("Piloto não encontrado");
        
        const pilot = await res.json();
        
        document.getElementById('pilot-info').innerHTML = `
            <h2>Dados do Piloto</h2>
            <p><strong>Nome:</strong> ${pilot.name}</p>
            <p><strong>Documento:</strong> ${pilot.document}</p>
            <p><strong>Telefone:</strong> ${pilot.phone}</p>
            <p><strong>Email:</strong> ${pilot.email}</p>
        `;
    } catch (error) {
        console.error(error);
        document.getElementById('pilot-info').innerHTML = `<p style="color:red;">Erro ao buscar dados do piloto.</p>`;
    }
}

document.addEventListener("DOMContentLoaded", fetchPilot);




