export interface Api {
  useAuthorize(): (providerId: string) => void;
  useSession: UseQuery<Session>;
}

export interface Session {
  user: User | null;
}

export interface QueryResult<Data> {
  data?: Data;
  error?: Error;
  loading: boolean;
}

export interface UseQuery<Data> {
  (): QueryResult<Data>;
}

export interface User {
  id: string;
  name: string;
  email: string;
}
