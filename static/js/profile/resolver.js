export async function profile()  {
    let profile;
    const hostname = window.location.hostname;

    switch (hostname) {
    case "rvkc-production.up.railway.app":
        profile = "/static/js/profile/prod.json";
        break;
    case "'rvkc-qa.up.railway.app'":
        profile = "/static/js/profile/qa.json";
        break;
    default:
        profile = "/static/js/profile/local.json";
    }

    const response = await fetch(profile);
    if (!response.ok) {
        throw new Error("Erro ao carregar config local.json");
    }
    return await response.json();
} 