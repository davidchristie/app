import { getByText, render, screen } from "../../testing";
import { ErrorPage } from ".";

describe("ErrorPage", () => {
  const title = "Something went wrong";

  it("renders correctly", () => {
    const { container } = render(
      <ErrorPage error={new Error()} resetErrorBoundary={() => {}} />
    );
    const element = screen.getByTestId("ErrorPage");
    expect(getByText(element, title)).toBeDefined();
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });

  describe("when 'Retry' button is clicked", () => {
    it("reloads the page", () => {
      const resetErrorBoundary = jest.fn();
      render(
        <ErrorPage
          error={new Error()}
          resetErrorBoundary={resetErrorBoundary}
        />
      );
      const retryButton = screen.getByTestId("ErrorPage__retryButton");
      retryButton.click();
      expect(resetErrorBoundary).toBeCalledTimes(1);
    });
  });
});
