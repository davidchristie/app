import { Suspense } from "react";
import { ErrorBoundary } from "./components/ErrorBoundary";
import { Header } from "./components/Header";

export function App() {
  return (
    <div data-testid="App">
      <ErrorBoundary>
        <Suspense fallback={<div data-testid="App__loading">Loading...</div>}>
          <Header />
          <main>
            <p>
              <a href="/api/v1/auth/github/authorize">Sign in with GitHub</a>
            </p>
            <p>
              <a href="/api/v1/auth/signout">Sign out</a>
            </p>
          </main>
        </Suspense>
      </ErrorBoundary>
    </div>
  );
}
