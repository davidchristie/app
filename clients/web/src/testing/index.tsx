import { ReactElement, ReactNode } from "react";
import { render, RenderOptions } from "@testing-library/react";
import { MemoryRouter } from "react-router-dom";
import { SWRConfig } from "swr";
import { MockApiProvider } from "../api/mock";

interface DefaultWrapperProps {
  children?: ReactNode;
}

function DefaultWrapper({ children }: DefaultWrapperProps): JSX.Element {
  return (
    <SWRConfig value={{ dedupingInterval: 0 }}>
      <MemoryRouter>
        <MockApiProvider>{children}</MockApiProvider>
      </MemoryRouter>
    </SWRConfig>
  );
}

const customRender = (ui: ReactElement, options?: RenderOptions) =>
  render(ui, { wrapper: DefaultWrapper, ...options });

export * from "@testing-library/react";

export { customRender as render };
