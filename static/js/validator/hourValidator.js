
import { BaseValidator }   from "./baseValidator.js";

export class HourValidator extends BaseValidator {
    constructor(input, elementErrorID) {
      super(input, elementErrorID);
      if (!this.isValid) {
        return;
      }

    if (!/^([01]\d|2[0-3]):([0-5]\d)$/.test(this.value)) {
      this.set("O horário deve estar no formato HH:MM.");
      return;
    }
    
    this.clean();
  }
}