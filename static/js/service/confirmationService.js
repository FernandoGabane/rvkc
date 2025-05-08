import { ErrorResponse }        from "../error/errorResponse.js";
import { AbstractService } from "./abstractService.js";

export class ConfirmationServiceImpl extends AbstractService {

  async create(confirmationsData) {
    const response = await fetch(`${this.profile.service_url}/confirmations`, {
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


  async getByClubId(clubId) {
    const response = await fetch(`${this.profile.service_url}/confirmations?club_id=${clubId}`, {
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


  async getByAccountId(accountId) {
    const response = await fetch(`${this.profile.service_url}/confirmations?account_id=${accountId}`, {
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
}