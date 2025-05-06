import { BaseValidator }   from "./baseValidator.js";

export class CheckboxValidator {
  constructor(input, elementErrorID) {

    this.input   = input;
    this.errorEl = document.getElementById(elementErrorID);
    this.value   = this.input.value;
    this.isValid = false;

    if (!this.input.length) {
      return false;
    }

    const anySelected = Array.from(this.input).some(cb => cb.checked);

    if (!anySelected) {
        this.errorEl.textContent = "Selecione pelo menos um clube.";
        this.errorEl.style.display = "block";
        this.isValid = false;
        return;
    }

    
    this.errorEl.textContent = "";
    this.errorEl.style.display = "block";
    this.isValid = true;
  }
}

