
import { BaseValidator }   from "./baseValidator.js";

export class DateValidator extends BaseValidator {
  constructor(input, elementErrorID) {
    super(input, elementErrorID);
    if (!this.isValid) {
      return;
    }

    if (!/^(0[1-9]|[12][0-9]|3[01])\/(0[1-9]|1[0-2])\/\d{4}$/.test(this.value)) {
      this.set("A data deve estar no formato DD/MM/AAAA.");
      return ;
    }

    this.clean();
  }
}