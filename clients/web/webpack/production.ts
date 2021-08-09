import "webpack-dev-server";
import { merge } from "webpack-merge";
import baseConfig from "./base";

const config = merge(baseConfig, {
  devtool: "source-map",
  mode: "production",
});

export default config;
