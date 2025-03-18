document.addEventListener("DOMContentLoaded", function () {
    // Prevenir envio do formulário e recarregamento da página
    document.querySelector(".form").addEventListener("submit", function (event) {
        event.preventDefault();
    });

    // Evento de clique para o botão "Registrar"
    document.getElementById("register-id").addEventListener("click", async function (event) {
        event.preventDefault();

        // Pegando os valores dos campos
        let name = document.getElementById("name").value.trim();
        let documentId = document.getElementById("pilot-document-id").value.trim();
        let email = document.getElementById("email").value.trim();
        let phone = document.getElementById("phone").value.trim();

        // Verificar se todos os campos foram preenchidos
        if (!name || !documentId || !email || !phone) {
            alert("Por favor, preencha todos os campos.");
            return;
        }

        console.log("Enviando para API:", { name, documentId, email, phone });

        try {
            // Fazendo a requisição para o backend
            let response = await fetch('/pilots/'+documentId, {
                method: "GET"
            });

            if (!response.ok) {
                throw new Error("Erro ao recuperar piloto.");
            }

            // Convertendo a resposta para JSON
            let pilot = await response.json();
            console.log("Resposta da API:", pilot);

            // Atualizar os dados na tela com a resposta do backend
            document.getElementById("pilot-info").innerHTML = `
                <h2>Dados do Piloto Cadastrado</h2>
                <p><strong>Nome:</strong> ${pilot.name}</p>
                <p><strong>Documento:</strong> ${pilot.document}</p>
                <p><strong>E-mail:</strong> ${pilot.email}</p>
                <p><strong>Telefone:</strong> ${pilot.phone}</p>
            `;

        } catch (error) {
            console.error(error);
            document.getElementById("pilot-info").innerHTML = `<p style="color:red;">Erro ao encontrar piloto piloto.</p>`;
        }
    });
});
