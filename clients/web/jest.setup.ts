import "@testing-library/jest-dom";
import { cache } from "swr";

beforeEach(() => {
  cache.clear();
});
