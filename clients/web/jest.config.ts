import { Config } from "@jest/types";

const config: Config.InitialOptions = {
  collectCoverageFrom: ["<rootDir>/src/**/*.{ts,tsx}"],
  coverageThreshold: {
    global: {
      branches: 100,
      functions: 100,
      lines: 100,
      statements: 100,
    },
  },
  preset: "ts-jest",
  setupFilesAfterEnv: ["<rootDir>/jest.setup.ts"],
  testEnvironment: "jsdom",
};

export default config;
