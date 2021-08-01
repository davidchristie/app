import { getByTestId } from "@testing-library/dom";

describe("when module is loaded", () => {
  let container: HTMLElement;

  beforeEach(async () => {
    container = document.createElement("div");
    container.id = "root";
    document.body.appendChild(container);
    await import(".");
  });

  afterEach(() => {
    document.body.removeChild(container);
  });

  it("renders App", () => {
    expect(getByTestId(container, "App")).toBeDefined();
  });
});
