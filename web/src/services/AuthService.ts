import { Constants } from "@/common/constants";
import { LoginReq, SignUpReq } from "@/models/auth";

export default class AuthenService {
  constructor() {}

  async login(req: LoginReq) {
    try {
      const res = await fetch(
        Constants.API_URL + Constants.AUTH_ENDPOINT + "/login",
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify(req),
          headers: {
            Accept: "application/json, text/plain",
            "Content-Type": "application/json;charset=UTF-8",
          },
        }
      );

      return res.json();
    } catch (error) {
      throw error;
    }
  }

  async signUp(req: SignUpReq) {
    try {
      const res = await fetch(Constants.API_URL + Constants.USER_ENDPOINT, {
        method: "POST",
        mode: "cors",
        body: JSON.stringify(req),
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
