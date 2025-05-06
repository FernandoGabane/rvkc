document.addEventListener("DOMContentLoaded", () => {

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
});