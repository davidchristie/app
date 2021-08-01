import useSWR from "swr";

export interface Session {
  user: User | null;
}

export interface User {
  id: string;
  name: string;
  email: string;
}

export function useSession(): Session {
  const { data } = useSWR("/api/v1/auth/session", {
    suspense: true,
  });
  return data;
}
