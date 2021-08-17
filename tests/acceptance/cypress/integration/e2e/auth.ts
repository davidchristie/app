describe("Auth", () => {
  beforeEach(() => {
    cy.visit("/signin");
  });

  it("can sign in with GitHub", () => {
    cy.getTestId("AppBar__signInButton").should("be.visible").click();

    cy.getTestId("SignInPage__GitHubAuthorizeButton")
      .should("be.visible")
      .click();

    cy.getTestId("UserMenu").should("be.visible");
  });

  it("can sign in with Google", () => {
    cy.getTestId("AppBar__signInButton").should("be.visible").click();

    cy.getTestId("SignInPage__GoogleAuthorizeButton")
      .should("be.visible")
      .click();

    cy.getTestId("UserMenu").should("be.visible");
  });
});
