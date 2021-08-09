import HtmlWebpackPlugin from "html-webpack-plugin";
import path from "path";
import { Configuration as Config } from "webpack";
import "webpack-dev-server";

const config: Config = {
  devServer: {
    historyApiFallback: true,
    port: 3000,
    proxy: {
      "/api": "http://localhost:4000",
    },
  },
  devtool: "inline-source-map",
  entry: "./src/index.tsx",
  mode: "development",
  module: {
    rules: [
      {
        exclude: /node_modules/,
        test: /\.tsx?$/,
        use: "ts-loader",
      },
    ],
  },
  output: {
    path: path.resolve(__dirname, "..", "build"),
    publicPath: "/",
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: "./public/index.html",
    }),
  ],
  resolve: {
    extensions: [".js", ".ts", ".tsx"],
  },
};

export default config;
