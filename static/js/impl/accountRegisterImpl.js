import { AccountNameValidator } from "../validator/nameValidator.js";
import { DocumentValidator }    from "../validator/cpfValidator.js";
import { EmailValidator }       from "../validator/emailValidator.js";
import { PhoneValidator }       from "../validator/phoneValidator.js";
import { applyMaskCpf }         from "../mask/cpfMask.js";
import { applyMaskPhone }       from "../mask/phoneMask.js";
import { AccountServiceImpl }   from "../service/accountService.js";
import { openModal }            from "../component/modal.js";
import { ErrorResponse }        from "../error/errorResponse.js";

document.addEventListener("DOMContentLoaded", function () {
  const form = document.querySelector("form");

  const nameInput     = document.getElementById("account-name");
  const surnameInput  = document.getElementById("account-surname");
  const cpfInput      = document.getElementById("account-document");
  const emailInput    = document.getElementById("account-email");
  const phoneInput    = document.getElementById("account-phone");

  emailInput.setAttribute("type", "text");

  applyMaskCpf(cpfInput);
  applyMaskPhone(phoneInput);

  [nameInput, cpfInput, emailInput, phoneInput].forEach(input => {
    input.addEventListener("input", () => {
      const errorEl = input.parentElement.querySelector(".error-message");
      if (errorEl) {
        input.classList.remove("invalid");
        errorEl.textContent = "";
        errorEl.style.display = "none";
      }
    });
  });


  form.addEventListener("submit", async function (e) {
    e.preventDefault();
    let valid = true;

    if (!new AccountNameValidator(nameInput, "account-name-error").isValid)        valid = false;
    if (!new AccountNameValidator(surnameInput, "account-surname-error").isValid)  valid = false;
    if (!new DocumentValidator(cpfInput, "account-document-error").isValid)        valid = false;
    if (!new EmailValidator(emailInput, "account-email-error").isValid)            valid = false;
    if (!new PhoneValidator(phoneInput, "account-phone-error").isValid)            valid = false;

    if (valid) {
      const payload = {
        document: cpfInput.value.replace(/\D/g, ""),
        name: nameInput.value.trim() + " " + surnameInput.value.trim(),
        phone: phoneInput.value.replace(/\D/g, "").trim(),
        email: emailInput.value.trim()
      };

      const accountsService = await new AccountServiceImpl().init();
      const accountResponse = await accountsService.create(payload);
      if (accountResponse instanceof ErrorResponse) {
        openModal(accountResponse, false);
        return;
      }

      openModal("Piloto registrado com sucesso!",true, () => {
        form.reset();
      });
    }
  });
});
