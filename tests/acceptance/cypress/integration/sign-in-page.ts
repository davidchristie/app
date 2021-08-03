describe("Sign In Page", () => {
  beforeEach(() => {
    cy.visit("/signin");
  });

  it("has app bar", () => {
    cy.getTestId("AppBar").should("be.visible");
  });

  it("has page heading", () => {
    cy.getTestId("PageHeading").should("contain.text", "Sign In");
  });

  it("can sign in with GitHub", () => {
    cy.getTestId("AppBar__signInButton").should("be.visible").click();

    cy.getTestId("SignInPage__authorizeButton-github")
      .should("be.visible")
      .click();

    cy.getTestId("AppBar__signOutButton").should("be.visible");
  });
});
