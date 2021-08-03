import { MockApiProvider } from "../../api";
import { render, screen } from "../../testing";
import { signedInSession, signedOutSession } from "../../testing/data";
import { HomePage } from ".";

describe("HomePage", () => {
  it("renders correctly", () => {
    const { container } = render(
      <MockApiProvider data={{ session: signedInSession }}>
        <HomePage />
      </MockApiProvider>
    );
    const element = screen.getByTestId("HomePage");
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });
});
