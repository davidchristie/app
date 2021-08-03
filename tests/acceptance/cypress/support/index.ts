declare namespace Cypress {
  interface Chainable {
    getTestId(value: string): Chainable<Element>;
  }
}

Cypress.Commands.add("getTestId", (testId: string) => {
  return cy.get(`[data-testid="${testId}"]`);
});
