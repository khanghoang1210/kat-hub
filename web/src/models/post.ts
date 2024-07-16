import { User } from "./user";

export type Post = {
  id: number;
  textContent: string;
  author: User;
  createdAt: Date;
  updatedAt: Date;
};
