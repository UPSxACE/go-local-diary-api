const path = require("path");

module.exports = {
  mode: "production",
  entry: {
    common: "./common.js",
    index: "./index.js",
  },
  output: {
    path: path.resolve(__dirname, "../public/dist"),
    filename: "[name].js",
  },
};
