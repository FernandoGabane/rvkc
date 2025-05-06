import { BaseValidator }   from "./baseValidator.js";

export class DocumentValidator extends BaseValidator {

  constructor(input, elementErrorID) {
      super(input, elementErrorID);
      if (!this.isValid) {
        return;
      }

      if (this.value.replace(/\D/g, "").length !== 11) {
        this.set("Digite um CPF válido.");
        return;  
      }
      
      this.clean();
  }
}
