import { PageHeading } from ".";
import { getByTestId, render, screen } from "../../testing";

describe("PageHeading", () => {
  it("renders correctly", () => {
    const { container } = render(
      <PageHeading>
        <div data-testid="children" />
      </PageHeading>
    );
    const element = screen.getByTestId("PageHeading");
    expect(getByTestId(element, "children")).toBeDefined();
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });
});
