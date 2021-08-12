import { renderHook } from "@testing-library/react-hooks";
import { signedInSession } from "../testing/data";
import { MockApiProvider } from "../testing/mocks";
import { useApi } from "./hook";

describe("useApi", () => {
  it("returns data from the API context", () => {
    const { result: apiResult } = renderHook(() => useApi(), {
      wrapper: ({ children }) => (
        <MockApiProvider data={{ session: signedInSession }}>
          {children}
        </MockApiProvider>
      ),
    });
    const { useSession } = apiResult.current;
    const { result: sessionResult } = renderHook(() => useSession());
    expect(sessionResult.current).toEqual({
      data: signedInSession,
      loading: false,
    });
  });

  it("returns error from the API context", () => {
    const sessionError = new Error("Test error");
    const { result: apiResult } = renderHook(() => useApi(), {
      wrapper: ({ children }) => (
        <MockApiProvider data={{ session: sessionError }}>
          {children}
        </MockApiProvider>
      ),
    });
    const { useSession } = apiResult.current;
    const { result: sessionResult } = renderHook(() => useSession());
    expect(sessionResult.current).toEqual({
      error: sessionError,
      loading: false,
    });
  });

  it("throws error if API context is not found", () => {
    const { result } = renderHook(() => useApi());
    expect(result.error).toEqual(new Error("API context not found"));
  });
});
