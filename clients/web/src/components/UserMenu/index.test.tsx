import { render, screen } from "../../testing";
import { user } from "../../testing/data";
import { MockApiProvider } from "../../testing/mocks";
import { UserMenu } from ".";

describe("UserMenu", () => {
  it("renders correctly", () => {
    const { container } = render(<UserMenu user={user} />);
    const element = screen.getByTestId("UserMenu");
    expect(element).toBe(container.firstChild);
    expect(element).toMatchSnapshot();
  });

  describe("when icon button is clicked", () => {
    it("renders user menu", () => {
      const { container } = render(<UserMenu user={user} />);
      screen.getByTestId("UserMenu__iconButton").click();
      const element = screen.getByTestId("UserMenu__menu");
      expect(element).toBeDefined();
      expect(element).toMatchSnapshot();
    });

    describe("when sign out menu item is clicked", () => {
      it("signs out", () => {
        const signOut = jest.fn();
        render(
          <MockApiProvider data={{ signOut }}>
            <UserMenu user={user} />
          </MockApiProvider>
        );
        screen.getByTestId("UserMenu__iconButton").click();
        screen.getByTestId("UserMenu__signOut").click();
        expect(signOut).toBeCalledTimes(1);
      });
    });
  });
});
