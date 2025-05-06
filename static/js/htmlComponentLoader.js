export function loadHtmlComponent(path, containerId, callback) {
    const run = () => {
      fetch(path)
        .then(response => {
          if (!response.ok) throw new Error(`Erro ao carregar ${path}`);
          return response.text();
        })
        .then(html => {
          const container = document.getElementById(containerId);
          if (container) {
            container.innerHTML = html;
            if (typeof callback === "function") {
              // console.log("Executando callback...");
              callback();
            }
          } else {
            console.warn(`Elemento com ID "${containerId}" não encontrado.`);
          }
        })
        .catch(error => console.error("Erro ao carregar componente:", error));
    };
  
    if (document.readyState === 'loading') {
      document.addEventListener("DOMContentLoaded", run);
    } else {
      run(); // DOM ready
    }
  }
  