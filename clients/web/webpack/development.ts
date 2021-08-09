import "webpack-dev-server";
import { merge } from "webpack-merge";
import baseConfig from "./base";

const config = merge(baseConfig, {
  devServer: {
    historyApiFallback: true,
    port: 3000,
    proxy: {
      "/api": "http://localhost:4000",
    },
  },
  devtool: "inline-source-map",
  mode: "development",
});

export default config;
