export class DefaultError {
  constructor(error, parameterName, description) {
    this.error = error || null;
    this.parameterName = parameterName || null;
    this.description = description || "Erro desconhecido";
  }

  toString() {
    if (this.parameterName) {
      return `${this.parameterName}: ${this.description}`;
    }
    return this.description;
  }
}

export class ErrorResponse {
  constructor(errors) {
    this.error_response = errors || [];
  }

  static async fromResponse(response) {
    try {
      const body = await response.json();
      if (body?.error_response && Array.isArray(body.error_response)) {
        const errors = body.error_response.map(err =>
          new DefaultError(err.error, err.parameter_name, err.description)
        );
        return new ErrorResponse(errors);
      }
    } catch (e) {
      console.error("Erro ao processar o corpo da resposta:", e);
    }

    return new ErrorResponse([new DefaultError(null, null, "Erro desconhecido")]);
  }
}
