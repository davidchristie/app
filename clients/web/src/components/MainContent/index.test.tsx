import { MainContent } from ".";
import { getByTestId, render, screen } from "../../testing";

describe("MainContent", () => {
  it("renders correctly", () => {
    const { container } = render(
      <MainContent>
        <div data-testid="children" />
      </MainContent>
    );
    const element = screen.getByTestId("MainContent");
    expect(getByTestId(element, "children")).toBeDefined();
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });
});
