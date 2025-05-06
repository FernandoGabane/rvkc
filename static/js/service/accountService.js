import { ErrorResponse }      from "../error/errorResponse.js";

export async function create(accountData) {
  const response = await fetch("http://localhost:8080/accounts", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
    body: JSON.stringify(accountData)
  });

  if (!response.ok) {
    return await ErrorResponse.fromResponse(response);
  }

  return await response.json();
}


export async function simple(accountId) {
  const response = await fetch(`http://localhost:8080/accounts/${accountId}/simple`, {
    method: "GET",
    headers: {
      "Accept": "application/json"
    }
  });

  if (!response.ok) {
    return await ErrorResponse.fromResponse(response);
  }

  return await response.json();
}


export async function getAll() {
  const response = await fetch(`http://localhost:8080/accounts/simple`, {
    method: "GET",
    headers: {
      "Accept": "application/json"
    }
  });

  if (!response.ok) {
    return await ErrorResponse.fromResponse(response);
  }

  return await response.json();
}