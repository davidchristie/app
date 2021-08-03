describe("Home Page", () => {
  beforeEach(() => {
    cy.visit("/");
  });

  it("has app bar", () => {
    cy.getTestId("AppBar").should("be.visible");
  });

  it("has page heading", () => {
    cy.getTestId("PageHeading").should("contain.text", "Home");
  });
});
