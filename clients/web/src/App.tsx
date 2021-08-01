import { Suspense } from "react";
import { ErrorBoundary } from "./components/ErrorBoundary";
import { Header } from "./components/Header";

export function App() {
  return (
    <div data-testid="App">
      <ErrorBoundary>
        <Suspense fallback={<div data-testid="App__loading">Loading...</div>}>
          <Header />
        </Suspense>
      </ErrorBoundary>
    </div>
  );
}
