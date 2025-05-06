export function extractErrorMessages(errorBody) {
    if (!errorBody) return ["Erro ao registrar piloto"];
  
    if (errorBody.error_response && Array.isArray(errorBody.error_response)) {
      return errorBody.error_response.map(err => err.description || "Erro desconhecido");
    }
  
    return ["Erro ao registrar piloto"];
  }
  