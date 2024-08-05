import { Constants } from "@/common/constants";


export default class UserService {
  constructor() {}
  async getUserByID(id: number) {
    try {
      const res = await fetch(Constants.API_URL + Constants.USER_ENDPOINT +`/${id}`, {
        method: "GET",
        mode: "cors",
        headers: {
          Accept: "application/json, text/plain",
          "Content-Type": "application/json;charset=UTF-8",
        },
      });

      return res.json();
    } catch (error) {
      throw error;
    }
  }
}
