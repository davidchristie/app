describe("Sign In Page", () => {
  beforeEach(() => {
    cy.intercept("/api/v1/auth/session", {
      fixture: "session/signed-out.json",
    });

    cy.visit("/signin");
  });

  it("has app bar", () => {
    cy.getTestId("AppBar").should("be.visible");
  });

  it("has page heading", () => {
    cy.getTestId("PageHeading").should("contain.text", "Sign In");
  });
});
