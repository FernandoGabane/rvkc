export class BaseValidator {
  constructor(input, elementErrorID) {
    this.input   = input;
    this.errorEl = document.getElementById(elementErrorID);
    this.value   = this.input.value.trim();
    this.isValid = false;

    if (!this.input) {
      throw new Error("Elemento de input não encontrado.");
    }

    if (!this.errorEl) {
      throw new Error("Elemento erro não encontrado.");
    }


    this.validateNotEmpty()
  }

  validateNotEmpty() {
    const value = this.input.value.trim();
    if (!value) {
      this.set("Preencha o campo.")
      return;
    }

    this.clean()
  }

  set(customMessage) {
    this.input.classList.add("invalid");
    this.errorEl.textContent = customMessage;
    this.errorEl.style.display = "block";
    this.isValid = false;
  }

  clean() {
    this.input.classList.remove("invalid");
    this.errorEl.textContent = "";
    this.errorEl.style.display = "none";
    this.isValid = true;
  }
}
