import { Box, CssBaseline } from "@material-ui/core";
import { ErrorBoundary } from "react-error-boundary";
import { BrowserRouter } from "react-router-dom";
import { ApiProvider } from "./api";
import { ErrorPage } from "./pages/ErrorPage";
import { Routes } from "./Routes";

export function App() {
  return (
    <ErrorBoundary FallbackComponent={ErrorPage}>
      <Box data-testid="App">
        <CssBaseline />
        <ApiProvider>
          <BrowserRouter>
            <Routes />
          </BrowserRouter>
        </ApiProvider>
      </Box>
    </ErrorBoundary>
  );
}
