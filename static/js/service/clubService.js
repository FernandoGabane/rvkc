import { ErrorResponse }   from "../error/errorResponse.js";
import { AbstractService } from "./abstractService.js";

export class ClubServiceImpl extends AbstractService {
  
  async getAll() {
    const response = await fetch(`${this.profile.service_url}/clubs`, {
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

  async create(clubData) {
    const response = await fetch(`${this.profile.service_url}/clubs`, {
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
}
