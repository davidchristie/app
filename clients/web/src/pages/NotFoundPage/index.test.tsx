import { MockApiProvider } from "../../api";
import { render, screen } from "../../testing";
import { signedInSession } from "../../testing/data";
import { NotFoundPage } from ".";

describe("NotFoundPage", () => {
  it("renders correctly", () => {
    const { container } = render(
      <MockApiProvider data={{ session: signedInSession }}>
        <NotFoundPage />
      </MockApiProvider>
    );
    const element = screen.getByTestId("NotFoundPage");
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });
});
