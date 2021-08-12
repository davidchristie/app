import { getByTestId, render, screen } from "../../testing";
import { signedInSession, signedOutSession } from "../../testing/data";
import { MockApiProvider } from "../../testing/mocks";
import { AppBar } from ".";

describe("AppBar", () => {
  describe("loading", () => {
    it("renders correctly", () => {
      const { container } = render(<AppBar />);
      const element = screen.getByTestId("AppBar");
      expect(element).toBe(container.firstChild);
      expect(element).toMatchSnapshot();
    });
  });

  describe("signed in", () => {
    it("renders correctly", () => {
      const { container } = render(
        <MockApiProvider data={{ session: signedInSession }}>
          <AppBar />
        </MockApiProvider>
      );
      const element = screen.getByTestId("AppBar");
      expect(getByTestId(element, "UserMenu")).toBeDefined();
      expect(element).toBe(container.firstChild);
      expect(element).toMatchSnapshot();
    });
  });

  describe("signed out", () => {
    it("renders correctly", () => {
      const { container } = render(
        <MockApiProvider data={{ session: signedOutSession }}>
          <AppBar />
        </MockApiProvider>
      );
      const element = screen.getByTestId("AppBar");
      expect(getByTestId(element, "AppBar__signInButton")).toBeDefined();
      expect(element).toBe(container.firstChild);
      expect(element).toMatchSnapshot();
    });
  });
});
