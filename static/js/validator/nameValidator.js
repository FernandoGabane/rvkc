
import { BaseValidator }   from "./baseValidator.js";

export class AccountNameValidator extends BaseValidator {
  constructor(input, elementErrorID) {
    super(input, elementErrorID);
    if (!this.isValid) {
      return;
    }

    if (!/^[A-Za-z]{2,}$/.test(this.value)) {
      this.set("Informe um nome válido (apenas letras e minímo 2 letras).");
      return;  
    }
    
    this.clean();
  }
}  