import { Page } from ".";
import { getByTestId, render, screen } from "../../testing";

describe("Page", () => {
  it("renders correctly", () => {
    const { container } = render(
      <Page>
        <div data-testid="children" />
      </Page>
    );
    const element = screen.getByTestId("Page");
    expect(getByTestId(element, "children")).toBeDefined();
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });
});
