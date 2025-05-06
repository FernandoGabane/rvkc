import { ClubNameValidator } from "/static/js/validator/clubNameValidator.js";
import { DateValidator }     from "/static/js/validator/dateValidator.js";
import { HourValidator }     from "/static/js/validator/hourValidator.js";
import { OptionValidator }   from "/static/js/validator/optionValidator.js";
import { openModal }         from "/static/js/component/modal.js";
import { setOptions }        from "/static/js/component/option.js";
import { toIsoWithOffset }   from "/static/js/util/dateTimeConverter.js";
import { ErrorResponse }     from "/static/js/error/errorResponse.js";
import { getAll }            from "/static/js/service/accountService.js";
import { create }            from "/static/js/service/clubService.js";

document.addEventListener("DOMContentLoaded", async function () {
  const form = document.querySelector("form");

  const clubNameInput    = document.getElementById("club-name");
  const accountId        = document.getElementById("club-coach-list");
  const dateInput        = document.getElementById("date-input");
  const startAtInput     = document.getElementById("start-at");
  const endAtInput       = document.getElementById("end-at");
  
  dateInput.setAttribute("type", "text");
  startAtInput.setAttribute("type", "text");
  endAtInput.setAttribute("type", "text");


  [clubNameInput, dateInput, startAtInput, endAtInput].forEach(input => {
    input.addEventListener("input", () => {
      const errorEl = input.parentElement.querySelector(".error-message");
      if (errorEl) {
        input.classList.remove("invalid");
        errorEl.textContent = "";
        errorEl.style.display = "none";
      }
    });
  });

  const accountSimpleResponsList = await getAll();
  if (accountSimpleResponsList instanceof ErrorResponse) {
    openModal("Erro ao recuperar lista de coaches. Tente novamente mais tarde.", false);
    return;
  }
  setOptions("club-coach-list", accountSimpleResponsList);
  

  form.addEventListener("submit", async function (e) {
    e.preventDefault();
    let valid = true;

    if (!new ClubNameValidator(clubNameInput, "club-name-error").isValid)  valid = false;
    if (!new OptionValidator(accountId, "club-coach-error").isValid)       valid = false;
    if (!new DateValidator(dateInput, "date-error").isValid)               valid = false;
    if (!new HourValidator(startAtInput, "start-at-error").isValid)        valid = false;
    if (!new HourValidator(endAtInput, "end-at-error").isValid)            valid = false;

    if (valid) {
      const startAt = toIsoWithOffset(dateInput.value.trim(), startAtInput.value.trim());
      const endAt   = toIsoWithOffset(dateInput.value.trim(), endAtInput.value.trim());

      const payload = {
        name: clubNameInput.value.trim(),
        date: dateInput.value.trim(),
        start_at: startAt,
        end_at: endAt,
        account_id: accountId.value,
        slots: 30
      }

      const accountResponse = await create(payload);
      if (accountResponse instanceof ErrorResponse) {
        openModal(accountResponse, false);
        return;
      }

      openModal("Club registrado com sucesso!",true, () => {
        form.reset();
      });
    }
  });
});
