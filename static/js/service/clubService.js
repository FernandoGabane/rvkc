import { ErrorResponse }      from "../error/errorResponse.js";

export async function getAll() {
  const response = await fetch("http://localhost:8080/clubs", {
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

export async function create(clubData) {
  const response = await fetch("http://localhost:8080/clubs", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
    body: JSON.stringify(clubData)
  });

  if (!response.ok) {
    return await ErrorResponse.fromResponse(response);
  }

  return await response.json();
}
