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

  it("has GitHub sign in button", () => {
    cy.getTestId("SignInPage__GitHubAuthorizeButton").should(
      "contain.text",
      "Sign in with GitHub"
    );
  });

  it("has Google sign in button", () => {
    cy.getTestId("SignInPage__GoogleAuthorizeButton").should(
      "contain.text",
      "Sign in with Google"
    );
  });
});
