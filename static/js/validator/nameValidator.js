
import { BaseValidator }   from "./baseValidator.js";

export class AccountNameValidator extends BaseValidator {
  constructor(input, elementErrorID) {
    super(input, elementErrorID);
    if (!this.isValid) {
      return;
    }

    if (!/^[A-Za-z]{2,}(?: [A-Za-z]+)*$/.test(this.value.trim())) {
      this.set("Informe um nome válido (apenas letras e minímo 2 caracteres).");
      return;  
    }
    
    this.clean();
  }
}  