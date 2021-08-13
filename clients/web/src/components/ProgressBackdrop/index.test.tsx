import { render, screen } from "../../testing";
import { ProgressBackdrop } from ".";

describe("ProgressBackdrop", () => {
  describe("open", () => {
    it("renders correctly", () => {
      const { container } = render(<ProgressBackdrop open />);
      const element = screen.getByTestId("ProgressBackdrop");
      expect(element).toBe(container.firstChild);
      expect(element).toMatchSnapshot();
    });
  });

  describe("closed", () => {
    it("renders correctly", () => {
      const { container } = render(<ProgressBackdrop open={false} />);
      const element = screen.getByTestId("ProgressBackdrop");
      expect(element).toBe(container.firstChild);
      expect(element).toMatchSnapshot();
    });
  });
});
