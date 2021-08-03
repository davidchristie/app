import { renderHook } from "@testing-library/react-hooks";
import { useApi } from "./hook";
import { ApiProvider } from "./provider";
import { Api } from "./types";

describe("ApiProvider", () => {
  describe("useApi", () => {
    let api: Api;

    beforeEach(() => {
      const { result: apiResult } = renderHook(() => useApi(), {
        wrapper: ({ children }) => <ApiProvider>{children}</ApiProvider>,
      });
      api = apiResult.current;
    });

    describe("useAuthorize", () => {
      const location = window.location;

      beforeEach(() => {
        Object.defineProperty(window, "location", {
          value: {
            href: "/",
          },
        });
      });

      afterAll(() => {
        Object.defineProperty(window, "location", {
          value: location,
        });
      });

      it("redirects to /api/v1/auth/{providerId}/authorize", () => {
        const { result: authorizeResult } = renderHook(
          () => api.useAuthorize(),
          {
            wrapper: ({ children }) => <ApiProvider>{children}</ApiProvider>,
          }
        );
        const providerId = "github";
        authorizeResult.current(providerId);
        expect(window.location.href).toBe(
          `/api/v1/auth/${providerId}/authorize`
        );
      });
    });
  });
});
