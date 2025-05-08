export function setOptions(dropdownElementId, data) {
  const clubSelect = document.getElementById(dropdownElementId);
  if (!clubSelect) return;

  data.forEach(d => {
    const option = document.createElement("option");
    option.value = d.id;
    option.textContent = d.name;
    clubSelect.appendChild(option);
  });
}
