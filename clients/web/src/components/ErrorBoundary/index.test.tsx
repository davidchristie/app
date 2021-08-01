import { render, screen } from "@testing-library/react";
import ErrorBoundary from ".";

describe("ErrorBoundary", () => {
  describe("when no error is thrown", () => {
    it("renders correctly", () => {
      function Child() {
        return <div data-testid="Child" />;
      }
      const { container } = render(
        <ErrorBoundary>
          <Child />
        </ErrorBoundary>
      );
      const child = screen.getByTestId("Child");
      expect(child).toBeDefined();
      expect(child).toBe(container.firstChild);
    });
  });

  describe("when error is thrown", () => {
    const message = "Something went wrong";
    const error = new Error(message);

    let consoleError: jest.SpyInstance;

    function ThrowError(): JSX.Element {
      throw error;
    }

    beforeEach(() => {
      consoleError = jest.spyOn(console, "error");
      consoleError.mockImplementation(() => {});
    });

    afterEach(() => {
      consoleError.mockRestore();
    });

    it("renders correctly", () => {
      const { container } = render(
        <ErrorBoundary>
          <ThrowError />
        </ErrorBoundary>
      );
      expect(screen.getByText(message)).toBeDefined();
      expect(container.firstChild).toMatchSnapshot();
    });

    it("calls onError handler", () => {
      const onError = jest.fn();
      render(
        <ErrorBoundary onError={onError}>
          <ThrowError />
        </ErrorBoundary>
      );
      expect(onError).toBeCalledTimes(1);
      expect(onError).toBeCalledWith(error, expect.any(Object));
    });

    it("logs error to console", () => {
      render(
        <ErrorBoundary>
          <ThrowError />
        </ErrorBoundary>
      );
      expect(consoleError).toBeCalledTimes(2);
      expect(consoleError.mock.calls[0][0]).toContain(
        `Error: Uncaught [Error: ${message}]`
      );
    });
  });
});
