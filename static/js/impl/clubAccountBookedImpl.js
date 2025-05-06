import { renderTable }              from "/static/js/component/table.js";
import { getAll }                   from "/static/js/service/clubService.js";
import { getByClubId }              from "/static/js/service/confirmationService.js";
import { openModal }                from "/static/js/component/modal.js";
import { ErrorResponse }            from "/static/js/error/errorResponse.js";
import { sortClubsByStartDateDesc } from "/static/js/util/sortList.js";


document.addEventListener("DOMContentLoaded", async function () {
  const form                 = document.querySelector("form");
  const clubsTable           = document.getElementById("clubs-table-body");
  const accountsBookedTable  = document.getElementById("accounts-booked-table-body");
  // const recordCountClub      = document.getElementById("record-count-clubs");
  const recordCountAccounts  = document.getElementById("record-count-accounts");

  const clubsResponse = await getAll();
  if (clubsResponse instanceof ErrorResponse) {
    openModal("Erro ao carregar clubes. Tente novamente mais tarde!", false);
    return;
  }

  const sortedClubs = sortClubsByStartDateDesc(clubsResponse);
  renderTable(sortedClubs, clubsTable, {
    type: "clubs",
    inputType: "radio",
    inputName: "clubs"
  });
  // recordCountClub.textContent = `Total de Clubs Abertos: ${clubsResponse.length}`;


  const clubInputs = form.querySelectorAll('input[name="clubs"]');
  clubInputs.forEach(input => {
    input.addEventListener("change", async function () {
      const clubId = input.value;
      const accountsResponse = await getByClubId(clubId);

      if (accountsResponse instanceof ErrorResponse) {
        openModal("Erro ao carregar contas. Tente novamente mais tarde!", false);
        return;
      }

      renderTable(accountsResponse, accountsBookedTable, {
        type: "club-account-booked",
      });
      recordCountAccounts.textContent = `Total de Pilotos Confirmados: ${accountsResponse.length}`;
    });
  });


});

