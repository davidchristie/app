import { ReactNode } from "react";
import useSWR from "swr";
import { authorize } from "./actions";
import { ApiContext } from "./context";
import { Api, UseQuery } from "./types";

export interface ApiProviderProps {
  children?: ReactNode;
}

function createQueryHook<Data>(url: string): UseQuery<Data> {
  return () => {
    const { data, error } = useSWR(url);
    const loading = !error && !data;
    return { data, error, loading };
  };
}

export function ApiProvider({ children }: ApiProviderProps): JSX.Element {
  const api: Api = {
    useAuthorize: () => authorize,
    useSession: createQueryHook("/api/v1/auth/session"),
  };
  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
}
