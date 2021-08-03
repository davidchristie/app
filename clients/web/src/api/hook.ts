import { useContext } from "react";
import { ApiContext } from "./context";
import { Api } from "./types";

export function useApi(): Api {
  const api = useContext(ApiContext);
  if (!api) {
    throw new Error("API context not found");
  }
  return api;
}
