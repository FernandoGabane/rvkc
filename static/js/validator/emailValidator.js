import { BaseValidator }   from "./baseValidator.js";

export class EmailValidator extends BaseValidator {

  constructor(input, elementErrorID) {
    super(input, elementErrorID);
    if (!this.isValid) {
      return;
    }

    if (!/.+@.+\..+/.test(this.value)) {
      this.set("E-mail inválido.");
      return;  
    }
    
    this.clean();
  }
}