import "@testing-library/jest-dom";
import { cache } from "swr";
import "whatwg-fetch";
import { server } from "./src/mocks/server";

beforeAll(() => server.listen());

beforeEach(() => {
  cache.clear();
});

afterEach(() => server.resetHandlers());

afterAll(() => server.close());
