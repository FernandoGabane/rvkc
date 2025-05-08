document.addEventListener("DOMContentLoaded", () => {
  const content = document.getElementById("app-content");

  document.body.addEventListener("click", async (e) => {
    const link = e.target.closest("[data-load]");
    if (!link) return;

    e.preventDefault();
    e.stopPropagation();

    const htmlPath = link.dataset.load;
    const scriptAttr = link.dataset.script || null;
    const styleAttr = link.dataset.style || null;
    const initAttr = link.dataset.init || null;

    try {
      const response = await fetch(htmlPath);
      if (!response.ok) throw new Error("Failed to load: " + htmlPath);
      const html = await response.text();
      content.innerHTML = html;

      document.body.classList.remove("__move");
      window.history.pushState(null, "", "/");

      // ✅ Inject styles
      if (styleAttr) {
        injectStyles(styleAttr);
      } else {
        injectStyles(inferStylePath(htmlPath));
      }

      // ✅ Load and init scripts
      const scriptPaths = (scriptAttr || inferScriptPath(htmlPath)).split(",").map(s => s.trim());
      const initFns = (initAttr || "").split(",").map(f => f.trim());

      for (let i = 0; i < scriptPaths.length; i++) {
        const scriptPath = scriptPaths[i];
        const initFn = initFns[i] || null;
        await importModuleAndInit(scriptPath, initFn);
      }

    } catch (err) {
      content.innerHTML = `<p>Failed to load content</p>`;
      console.error(err);
    }
  });
});

function injectStyles(styleList) {
  const styles = styleList.split(",").map(s => s.trim());
  styles.forEach(cssPath => {
    if (document.querySelector(`link[href="${cssPath}"]`)) return;
    const link = document.createElement("link");
    link.rel = "stylesheet";
    link.href = cssPath;
    document.head.appendChild(link);
  });
}

function inferStylePath(htmlPath) {
  const name = htmlPath.split("/").pop().replace(".html", ".css");
  return `/static/css/${name}`;
}

function inferScriptPath(htmlPath) {
  const name = htmlPath.split("/").pop().replace(".html", "Impl.js");
  return `/static/js/impl/${name}`;
}

function getInitFunctionName(scriptPath) {
  const base = scriptPath.split("/").pop().replace("Impl.js", "");
  const [first, ...rest] = base.split(/[-_]/);
  const capitalized = [first, ...rest.map(word => word.charAt(0).toUpperCase() + word.slice(1))].join("");
  return `init${capitalized.charAt(0).toUpperCase() + capitalized.slice(1)}`;
}

async function importModuleAndInit(scriptPath, initFunctionName = null) {
  try {
    const module = await import(scriptPath);
    const functionName = initFunctionName || getInitFunctionName(scriptPath);

    if (initFunctionName && typeof module[functionName] === "function") {
      module[functionName]();
    } else if (initFunctionName) {
      console.warn(`⚠️ Function ${functionName}() not found in ${scriptPath}`);
    }
  } catch (err) {
    console.warn(`⚠️ Failed to import module: ${scriptPath}`, err);
  }
}
