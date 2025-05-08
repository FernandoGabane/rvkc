export function renderTable(dataList, tableBodyElement, options = {}) {
  tableBodyElement.innerHTML = "";

  let createRowFn;

  switch (options.type) {
    case "clubs":
      createRowFn = (data) => createClubRow(data, options);
      break;
    case "accounts":
      createRowFn = (data) => createAccountRow(data, options);
      break;
    case "accounts-rules":
      createRowFn = (data) => createAccountRuleRow(data, options);
      break;      
    case "club-account-booked":
      createRowFn = (data) => createClubAccountBookedRuleRow(data, options);
      break;      
    default:
      throw new Error(`Tipo de linha desconhecido: ${options.type}`);
  }

  dataList.forEach(data => {
    const row = createRowFn(data);
    tableBodyElement.appendChild(row);
  });
}


function createClubRow(data, options) {
  const startDate = new Date(data.start_at);
  const day = startDate.toLocaleDateString('pt-BR');
  const hour = startDate.toLocaleTimeString('pt-BR', {
    hour: '2-digit',
    minute: '2-digit'
  });

  const inputCell = createInputCell(data, options);

  const row = document.createElement('tr');
  row.innerHTML = `
    ${inputCell}
    <td>${day}</td>
    <td>${data.name}</td>
    <td>${data.weekday}</td>
    <td>${hour}</td>
  `;
  return row;
}


function createAccountRow(data, options) {
  const inputCell = createInputCell(data, options);

  const row = document.createElement('tr');
  row.innerHTML = `
    ${inputCell}
    <td>${data.name}</td>
  `;
  return row;
}


function createAccountRuleRow(data, options) {
  const inputCell = createInputCell(data, options);

  const row = document.createElement('tr');
  row.innerHTML = `
    ${inputCell}
    <td>${data.name}</td>
    <td>${data.roles}</td>
  `;
  return row;
}


function createClubAccountBookedRuleRow(data, options) {
  const inputCell = createInputCell(data, options);

  const row = document.createElement('tr');
  row.innerHTML = `
    ${inputCell}
    <td>${data.account.name}</td>
  `;
  return row;
}


function createInputCell(data, options) {
  if (!options?.inputType || !options?.inputName) return "";

  switch (options.inputType) {
    case "checkbox":
    case "radio":
      return `<td><input type="${options.inputType}" name="${options.inputName}" value="${data.id}"></td>`;
    default:
      console.warn(`Tipo de input inválido: ${options.inputType}`);
      return "";
  }
}