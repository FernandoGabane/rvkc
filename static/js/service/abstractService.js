import { profile }  from "../profile/resolver.js";

export class AbstractService {
  async init() {
    this.profile = await profile();
    return this;
  }
}
  