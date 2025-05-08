export function applyMaskPhone(input) {
    input.addEventListener("input", function () {
      const valor = this.value.replace(/\D/g, "").slice(0, 11);
      let formatado = "";
  
      if (valor.length <= 10) {
        formatado = valor.replace(/(\d{0,2})(\d{0,4})(\d{0,4})/, (_, ddd, parte1, parte2) => {
          return (ddd ? `(${ddd}` : "") + (ddd.length === 2 ? ") " : "") + (parte1 || "") + (parte1?.length === 4 && parte2 ? "-" + parte2 : "");
        });
      } else {
        formatado = valor.replace(/(\d{0,2})(\d{0,5})(\d{0,4})/, (_, ddd, parte1, parte2) => {
          return (ddd ? `(${ddd}` : "") + (ddd.length === 2 ? ") " : "") + (parte1 || "") + (parte1?.length >= 5 && parte2 ? "-" + parte2 : "");
        });
      }
  
      this.value = formatado;
    });
  }
  
  