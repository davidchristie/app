import { ReactNode } from "react";
import { ApiContext } from "../../../api/context";
import { Api, Session, UseQuery } from "../../../api/types";

export interface MockApiData {
  authorize?: () => void;
  session?: Session | Error;
  signOut?: () => void;
}

export interface MockApiProviderProps {
  children?: ReactNode;
  data?: MockApiData;
}

function createQueryHook<Data>(
  value: Data | Error | undefined
): UseQuery<Data> {
  return () => {
    if (!value) {
      return {
        loading: true,
      };
    } else if (value instanceof Error) {
      return {
        error: value,
        loading: false,
      };
    }
    return {
      data: value,
      loading: false,
    };
  };
}

export function MockApiProvider({
  children,
  data = {},
}: MockApiProviderProps): JSX.Element {
  const api: Api = {
    useAuthorize: () => data.authorize ?? jest.fn(),
    useSession: createQueryHook(data.session),
    useSignOut: () => data.signOut ?? jest.fn(),
  };
  return <ApiContext.Provider value={api}>{children}</ApiContext.Provider>;
}
