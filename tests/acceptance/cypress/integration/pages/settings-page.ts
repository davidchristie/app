describe("Settings Page", () => {
  beforeEach(() => {
    cy.intercept("/api/v1/auth/session", { fixture: "session/signed-in.json" });

    cy.visit("/settings");
  });

  it("has app bar", () => {
    cy.getTestId("AppBar").should("be.visible");
  });

  it("has page heading", () => {
    cy.getTestId("PageHeading").should("contain.text", "Settings");
  });
});
