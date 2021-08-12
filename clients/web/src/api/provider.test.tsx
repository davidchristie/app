import { renderHook } from "@testing-library/react-hooks";
import useSWR from "swr";
import { useApi } from "./hook";
import { ApiProvider } from "./provider";
import { Api } from "./types";

jest.mock("swr");

describe("ApiProvider", () => {
  describe("useApi", () => {
    const location = window.location;

    let api: Api;

    beforeEach(() => {
      jest.resetAllMocks();
      Object.defineProperty(window, "location", {
        value: {
          href: "/",
        },
      });
      const { result: apiResult } = renderHook(() => useApi(), {
        wrapper: ({ children }) => <ApiProvider>{children}</ApiProvider>,
      });
      api = apiResult.current;
    });

    afterAll(() => {
      Object.defineProperty(window, "location", {
        value: location,
      });
    });

    describe("useAuthorize", () => {
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

    describe("useSession", () => {
      describe("loading", () => {
        it("returns correct result", () => {
          (useSWR as jest.Mock).mockReturnValueOnce({
            data: undefined,
            error: undefined,
            loading: true,
          });
          const { result: sessionResult } = renderHook(() => api.useSession(), {
            wrapper: ({ children }) => <ApiProvider>{children}</ApiProvider>,
          });
          expect(sessionResult.current).toEqual({
            data: undefined,
            error: undefined,
            loading: true,
          });
        });
      });

      describe("data", () => {
        it("returns correct result", () => {
          (useSWR as jest.Mock).mockReturnValueOnce({
            data: {
              user: {
                avatar_url: "https://via.placeholder.com/150",
                id: "97406d59-7a49-4f1e-bb79-aba34cfcb405",
                name: "Test User",
                email: "test_user@email.com",
              },
            },
            error: undefined,
            loading: false,
          });
          const { result: sessionResult } = renderHook(() => api.useSession(), {
            wrapper: ({ children }) => <ApiProvider>{children}</ApiProvider>,
          });
          expect(sessionResult.current).toEqual({
            data: {
              user: {
                avatarUrl: "https://via.placeholder.com/150",
                id: "97406d59-7a49-4f1e-bb79-aba34cfcb405",
                name: "Test User",
                email: "test_user@email.com",
              },
            },
            error: undefined,
            loading: false,
          });
        });
      });
    });

    describe("useSignOut", () => {
      it("redirects to /api/v1/auth/signout", () => {
        const { result: signOutResult } = renderHook(() => api.useSignOut(), {
          wrapper: ({ children }) => <ApiProvider>{children}</ApiProvider>,
        });
        signOutResult.current();
        expect(window.location.href).toBe("/api/v1/auth/signout");
      });
    });
  });
});
