export function initDatepickr() {
  if (window.__flatpickr_initialized__) return;

  // Verifica se a lib está disponível
  if (typeof flatpickr !== "function") {
    console.warn("⚠️ flatpickr is not loaded.");
    return;
  }

  flatpickr.localize(flatpickr.l10ns.pt);

  flatpickr("#date-input", {
    dateFormat: "d/m/Y"
  });

  flatpickr("#start-at", {
    enableTime: true,
    noCalendar: true,
    dateFormat: "H:i",
    time_24hr: true
  });

  flatpickr("#end-at", {
    enableTime: true,
    noCalendar: true,
    dateFormat: "H:i",
    time_24hr: true
  });

  console.log("✅ Flatpickr initialized");
  window.__flatpickr_initialized__ = true;
}
