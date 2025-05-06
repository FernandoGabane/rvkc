import { ErrorResponse } from "../error/errorResponse.js"; // ajuste o caminho conforme necessário

export function openModal(messages, success, onClose) {
  const modal = document.createElement("div");
  modal.classList.add("custom-modal-overlay");

  const box = document.createElement("div");
  box.classList.add("custom-modal-box");

  const p = document.createElement("div");
  p.classList.add("custom-modal-message");

  const ul = document.createElement("ul");

  // Verifica se é uma instância de ErrorResponse
  if (messages instanceof ErrorResponse) {
    (messages.error_response || []).forEach(err => {
      const li = document.createElement("li");
      li.textContent = err.description || "Erro desconhecido";
      ul.appendChild(li);
    });
  } else if (Array.isArray(messages)) {
    // caso ainda seja um array comum de mensagens
    messages.forEach(msg => {
      const li = document.createElement("li");
      li.textContent = typeof msg === "string" ? msg : "Erro desconhecido";
      ul.appendChild(li);
    });
  } else {
    const li = document.createElement("li");
    li.textContent = typeof messages === "string" ? messages : "Erro desconhecido";
    ul.appendChild(li);
  }

  p.appendChild(ul);

  const button = document.createElement("button");
  button.classList.add("custom-modal-button");
  button.innerText = "OK";
  button.style.backgroundColor = success ? "#28a745" : "#dc3545";

  button.addEventListener("click", () => {
    document.body.removeChild(modal);
    if (typeof onClose === "function") {
      onClose();
    }
  });

  box.appendChild(p);
  box.appendChild(button);
  modal.appendChild(box);
  document.body.appendChild(modal);
}
