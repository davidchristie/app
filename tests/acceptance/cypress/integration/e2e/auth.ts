describe("Auth", () => {
  beforeEach(() => {
    cy.visit("/signin");
  });

  it("can sign in with GitHub", () => {
    cy.getTestId("AppBar__signInButton").should("be.visible").click();

    cy.getTestId("SignInPage__authorizeButton-github")
      .should("be.visible")
      .click();

    cy.getTestId("UserMenu").should("be.visible");
  });
});
