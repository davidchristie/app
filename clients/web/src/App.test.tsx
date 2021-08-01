import {
  render,
  screen,
  waitForElementToBeRemoved,
} from "@testing-library/react";
import { App } from "./App";

describe("App", () => {
  describe("while loading", () => {
    it("renders correctly", () => {
      const { container } = render(<App />);
      expect(screen.getByTestId("App__loading")).toBeDefined();
      expect(container.firstChild).toMatchSnapshot();
    });
  });

  describe("once loading is complete", () => {
    it("renders correctly", async () => {
      const { container } = render(<App />);
      await waitForElementToBeRemoved(() =>
        screen.queryByTestId("App__loading")
      );
      expect(container.firstChild).toMatchSnapshot();
    });
  });
});
