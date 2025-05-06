import { renderTable } from "../component/table.js";
import { getAll }              from "../service/accountService.js";
import { openModal }           from "../component/modal.js";
import { ErrorResponse }       from "../error/errorResponse.js";

document.addEventListener("DOMContentLoaded", async function () {
  const form = document.querySelector("form");

  const tableBody = document.getElementById("list-account-table-body");
  const recordCount = document.getElementById("record-count")


  const accountsList = await getAll();
  if (accountsList instanceof ErrorResponse) {
    openModal("Erro ao carregar liste Pilotos. Tente novamente mais tarde!", false);
    return;
  }

  renderTable(accountsList, tableBody, {
    type: "accounts"
  });

  recordCount.textContent = `Total de Pilotos: ${accountsList.length}`;

});