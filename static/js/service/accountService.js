import { ErrorResponse }   from "../error/errorResponse.js";
import { AbstractService } from "./abstractService.js";

export class AccountServiceImpl extends AbstractService {

  async create(accountData) {
    const response = await fetch(`${this.profile.service_url}/accounts`, {
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
  
  
  async simple(accountId) {
    const response = await fetch(`${this.profile.service_url}/accounts/${accountId}/simple`, {
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


  async getAll() {
    const response = await fetch(`${this.profile.service_url}/accounts/simple`, {
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