export function applyMaskCpf(input) {
    input.addEventListener("input", function () {
      let valor = this.value.replace(/\D/g, "");
      valor = valor.slice(0, 11);
      valor = valor.replace(/(\d{3})(\d)/, "$1.$2");
      valor = valor.replace(/(\d{3})(\d)/, "$1.$2");
      valor = valor.replace(/(\d{3})(\d{1,2})$/, "$1-$2");
      this.value = valor;
    });
  }
