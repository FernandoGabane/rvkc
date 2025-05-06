import { ErrorResponse }        from "../error/errorResponse.js";

export async function create(confirmationsData) {
  const response = await fetch("http://localhost:8080/confirmations", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Accept": "application/json"
    },
    body: JSON.stringify(confirmationsData)
  });

  if (!response.ok) {
    return await ErrorResponse.fromResponse(response);
  }

  return await response.json();
}


export async function getByClubId(clubId) {
  const response = await fetch(`http://localhost:8080/confirmations?club_id=${clubId}`, {
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


export async function getByAccountId(accountId) {
  const response = await fetch(`http://localhost:8080/confirmations?account_id=${accountId}`, {
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