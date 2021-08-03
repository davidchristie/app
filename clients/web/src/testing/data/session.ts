import { Session } from "../../api";
import { user } from "./user";

export const signedInSession: Session = {
  user,
};

export const signedOutSession: Session = {
  user: null,
};
