import { menuActions } from './menuAction.js';
import { loadHtmlComponent } from './htmlComponentLoader.js';


const roleFromBody = document.body.dataset.role;
const role         = roleFromBody || localStorage.getItem("userRole") || "account";

if (role === "admin") {
  loadHtmlComponent("/static/components/menuAdmin.html", "menu-container", menuActions);
} else {
  loadHtmlComponent("/static/components/menuAccount.html", "menu-container", menuActions);
}
