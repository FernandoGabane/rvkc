import { BaseValidator }   from "./baseValidator.js";

export class OptionValidator extends BaseValidator {
    constructor(input, elementErrorID) {
      super(input, elementErrorID);
      if (!this.isValid) {
        return;
      } 

      if (this.value.toLowerCase().includes("selecione") || this.value.toLowerCase().includes("select")) {
        this.set("Selecione uma opção válida.");
      }

      this.clean();
  }
}