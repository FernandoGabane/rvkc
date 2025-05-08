export function menuActions() {
  // Submenus com seta
  document.querySelectorAll('.link-nav-arrow').forEach(link => {
    link.addEventListener('click', function (e) {
      e.preventDefault();

      const submenu = this.nextElementSibling;
      if (submenu?.classList.contains('list-nav-second')) {
        submenu.classList.toggle('hide');
        this.classList.toggle('arrow');
      }
    });
  });

  // Menu hambúrguer
  const iconMenu = document.querySelector(".icon-menu");
  iconMenu?.addEventListener("click", (e) => {
    e.stopPropagation(); // <-- impede que o clique feche o menu
    document.body.classList.toggle("__move");
  });

  // Clique fora do menu fecha o menu
  const content = document.querySelector(".content");
  content?.addEventListener("click", () => {
    document.body.classList.remove("__move");
  });

  // Ativar link selecionado na navegação
  const navLinks = document.querySelectorAll(".link-nav");
  navLinks.forEach(link => {
    link.addEventListener("click", () => {
      navLinks.forEach(l => l.classList.remove("active"));
      link.classList.add("active");
    });
  });
}
