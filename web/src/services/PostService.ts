import { Constants } from "@/common/constants";

const POST_ENDPOINT = "/posts";
export class PostService {
    constructor(){
       
    }
    async getPosts(){
        try {
            const res = await fetch(Constants.API_URL + POST_ENDPOINT, {
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