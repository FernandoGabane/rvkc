export async function profile()  {
    let profile;

    switch (window.PROFILE) {
    case "prod":
        profile = "/static/js/profile/prod.json";
        break;
    case "qa":
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
