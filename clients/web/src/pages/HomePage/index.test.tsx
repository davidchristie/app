import { render, screen } from "../../testing";
import { signedInSession } from "../../testing/data";
import { MockApiProvider } from "../../testing/mocks";
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
