import { StaticRouter } from "react-router-dom";
import { render, screen } from "./testing";
import { signedInSession, signedOutSession } from "./testing/data";
import { MockApiProvider } from "./testing/mocks";
import { Routes } from "./Routes";

describe("Routes", () => {
  describe("loading", () => {
    it("renders correctly", () => {
      const { container } = render(<Routes />);
      expect(container.firstChild).toMatchSnapshot();
    });
  });

  describe("signed in", () => {
    const routes: { expectedTestId: string; url: string }[] = [
      {
        expectedTestId: "HomePage",
        url: "/",
      },
      {
        expectedTestId: "SettingsPage",
        url: "/settings",
      },
    ];

    routes.forEach(({ expectedTestId, url }) => {
      it(url, () => {
        render(
          <StaticRouter location={url}>
            <MockApiProvider data={{ session: signedInSession }}>
              <Routes />
            </MockApiProvider>
          </StaticRouter>
        );
        expect(screen.getAllByTestId(expectedTestId)).toBeDefined();
      });
    });
  });

  describe("signed out", () => {
    const routes: { expectedTestId: string; url: string }[] = [
      {
        expectedTestId: "HomePage",
        url: "/",
      },
      {
        expectedTestId: "SignInPage",
        url: "/signin",
      },
    ];

    routes.forEach(({ expectedTestId, url }) => {
      it(url, () => {
        render(
          <StaticRouter location={url}>
            <MockApiProvider data={{ session: signedOutSession }}>
              <Routes />
            </MockApiProvider>
          </StaticRouter>
        );
        expect(screen.getAllByTestId(expectedTestId)).toBeDefined();
      });
    });
  });
});
