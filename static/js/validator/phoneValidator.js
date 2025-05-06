import { BaseValidator }   from "./baseValidator.js";

export class PhoneValidator extends BaseValidator {
  constructor(input, elementErrorID) {
    super(input, elementErrorID);
    if (!this.isValid) {
      return;
    }

    const phoneNumber = this.value.replace(/\D/g, "");
    
    if (phoneNumber.length < 10 || phoneNumber.length > 11) {
      this.set("Telefone deve ter DDD e de 8 a 9 dígitos.")
      return;  
    }
    
    this.clean();
  }
}