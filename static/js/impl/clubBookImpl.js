import { DocumentValidator }        from "/static/js/validator/cpfValidator.js";
import { CheckboxValidator }        from "/static/js/validator/checkboxValidator.js";
import { applyMaskCpf }             from "/static/js/mask/cpfMask.js";
import { renderTable }              from "/static/js/component/table.js";
import { getAll }                   from "/static/js/service/clubService.js";
import { openModal }                from "/static/js/component/modal.js";
import { simple }                   from "/static/js/service/accountService.js";
import { create }                   from "/static/js/service/confirmationService.js";
import { ErrorResponse }            from "/static/js/error/errorResponse.js";
import { sortClubsByStartDateDesc } from "/static/js/util/sortList.js";

document.addEventListener("DOMContentLoaded", async function () {
  const form       = document.querySelector("form");
  const tableBody  = document.getElementById("clubs-table-body");
  const cpfInput   = document.getElementById("account-document");

  applyMaskCpf(cpfInput);

  [cpfInput, tableBody].forEach(input => {
    input.addEventListener("input", () => {
      const errorEl = input.parentElement.querySelector(".error-message");
      if (errorEl) {
        input.classList.remove("invalid");
        errorEl.textContent = "";
        errorEl.style.display = "none";
      }
    });
  });

  const clubsResponse = await getAll();
  if (clubsResponse instanceof ErrorResponse) {
    openModal("Erro ao carregar clubes. Tente novamente mais tarde!", false);
    return;
  }

  
  const sortedClubs = sortClubsByStartDateDesc(clubsResponse);
  renderTable(sortedClubs, tableBody, {
    type: "clubs",
    inputType: "checkbox",
    inputName: "clubs"
  });

  // renderClubsTable(clubsResponse, tableBody);

  form.addEventListener("submit", async function (e) {
    e.preventDefault();
    let valid = true;
    const checkboxes = form.querySelectorAll('input[type="checkbox"][name="clubs"]');

    if (!new DocumentValidator(cpfInput, "account-document-error").isValid)   valid = false;
    if (!new CheckboxValidator(checkboxes, "clubs-table-error").isValid)      valid = false;

    if (valid) {
      const accountResponse = await simple(cpfInput.value.replace(/\D/g, ""));

      if (accountResponse instanceof ErrorResponse) {
        openModal(accountResponse, false);
        return;
      }
      
      const selectedClubIds = Array.from(form.querySelectorAll('input[type="checkbox"][name="clubs"]:checked'))
        .map(cb => cb.value);

      const payload = {
        confirmations: selectedClubIds.map(clubId => ({
          club_id: clubId,
          account_id: accountResponse.id,
          status: "CONFIRMED"
        }))
      };
  
      const confirmationResponse = await create(payload);
      if (confirmationResponse instanceof ErrorResponse) {
        openModal(confirmationResponse, false);
        return;
      }

      openModal("Confirmações registradas com sucesso!", true, () => {
        form.reset();

        document.querySelectorAll('input[type="checkbox"][name="club"]').forEach(cb => {
          cb.setCustomValidity("");
        });
      });
    }
  });
});

