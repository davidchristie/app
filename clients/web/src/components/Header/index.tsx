import { useSession } from "../../hooks/useSession";

export function Header() {
  const session = useSession();
  return (
    <header data-testid="Header">
      <pre>{JSON.stringify(session, null, 2)}</pre>
    </header>
  );
}
