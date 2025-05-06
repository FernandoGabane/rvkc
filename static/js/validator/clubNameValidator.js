import { BaseValidator }   from "./baseValidator.js";

export class ClubNameValidator extends BaseValidator {
  constructor(input, elementErrorID) {
    super(input, elementErrorID);
    if (!this.isValid) {
      return;
    }

    const value = this.input.value.trim();
    if (value.length < 4) {
      this.set("O nome do clube deve ter pelo menos 4 caracteres.");
      return;  
    }
    
    this.clean();
  }
}
