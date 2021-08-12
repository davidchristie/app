import { render, screen } from "../../testing";
import { signedInSession } from "../../testing/data";
import { MockApiProvider } from "../../testing/mocks";
import { SignInPage } from ".";

describe("SignInPage", () => {
  it("renders correctly", () => {
    const { container } = render(
      <MockApiProvider data={{ session: signedInSession }}>
        <SignInPage />
      </MockApiProvider>
    );
    const element = screen.getByTestId("SignInPage");
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });

  it("can sign in with GitHub", () => {
    const authorize = jest.fn();
    render(
      <MockApiProvider data={{ authorize, session: signedInSession }}>
        <SignInPage />
      </MockApiProvider>
    );
    screen.getByTestId("SignInPage__authorizeButton-github").click();
    expect(authorize).toBeCalledTimes(1);
    expect(authorize).toBeCalledWith("github");
  });
});
