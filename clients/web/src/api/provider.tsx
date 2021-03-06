import camelCaseKeys from "camelcase-keys";
import { ReactNode } from "react";
import useSWR from "swr";
import { authorize, signOut } from "./actions";
import { ApiContext } from "./context";
import { Api, UseQuery } from "./types";

export interface ApiProviderProps {
  children?: ReactNode;
}

function createQueryHook<Data>(url: string): UseQuery<Data> {
  return () => {
    const { data, error } = useSWR(url);
    return {
      data: data ? camelCaseKeys(data, { deep: true }) : undefined,
      error,
      loading: !error && !data,
    };
  };
}

export function ApiProvider({ children }: ApiProviderProps): JSX.Element {
  const api: Api = {
    useAuthorize: () => authorize,
    useSession: createQueryHook("/api/v1/auth/session"),
    useSignOut: () => signOut,
  };
  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
}
