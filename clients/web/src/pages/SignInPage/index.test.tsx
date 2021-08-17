import { render, screen } from "../../testing";
import { signedOutSession } from "../../testing/data";
import { MockApiProvider } from "../../testing/mocks";
import { SignInPage } from ".";

describe("SignInPage", () => {
  it("renders correctly", () => {
    const { container } = render(
      <MockApiProvider data={{ session: signedOutSession }}>
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
      <MockApiProvider data={{ authorize, session: signedOutSession }}>
        <SignInPage />
      </MockApiProvider>
    );
    screen.getByTestId("SignInPage__GitHubAuthorizeButton").click();
    expect(authorize).toBeCalledTimes(1);
    expect(authorize).toBeCalledWith("github");
  });

  it("can sign in with Google", () => {
    const authorize = jest.fn();
    render(
      <MockApiProvider data={{ authorize, session: signedOutSession }}>
        <SignInPage />
      </MockApiProvider>
    );
    screen.getByTestId("SignInPage__GoogleAuthorizeButton").click();
    expect(authorize).toBeCalledTimes(1);
    expect(authorize).toBeCalledWith("google");
  });
});
