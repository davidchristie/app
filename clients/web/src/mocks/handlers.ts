import { RequestHandler, rest } from "msw";

export const handlers: RequestHandler[] = [
  rest.get("/api/v1/auth/session", (request, response, context) => {
    return response(
      context.status(200),
      context.body(
        JSON.stringify({
          user: {
            id: "97406d59-7a49-4f1e-bb79-aba34cfcb405",
            name: "Test User",
            email: "test_user@email.com",
          },
        })
      )
    );
  }),
];
