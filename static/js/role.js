const role = document.body.dataset.role;

if (role) {
  localStorage.setItem("userRole", role);
}
