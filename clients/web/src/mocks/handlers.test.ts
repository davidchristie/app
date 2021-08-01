interface Handler {
  method?: "DELETE" | "GET" | "POST";
  url: string;
}

const handlers: Handler[] = [
  {
    url: "/api/v1/auth/session",
  },
];

handlers.forEach(({ method = "GET", url }) => {
  describe(`${method} ${url}`, () => {
    it("responds correctly", async () => {
      const response = await fetch("/api/v1/auth/session");
      const body = await response.json();
      expect(response.ok).toBe(true);
      expect(response.status).toBe(200);
      expect(body).toMatchSnapshot();
    });
  });
});
