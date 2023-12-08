const path = require("path");
const glob = require("glob");

module.exports = {
  mode: "production",
  entry: () => {
    // NOTE: Please make sure you don't have any files with equal names, even in nested directories.
    const entryPoints = {};
    const directories = ["./"]; // make sure you finish each entry with slash

    directories.forEach((dir) => {
      const files = glob.sync(dir + "*.js"); // Adjust the pattern as needed

      files.forEach((file) => {
        if (file === "webpack.config.js") {
          return;
        }
        const entryName = path.basename(file, path.extname(file));
        entryPoints[entryName] = "./" + file.replace("\\", "/");
      });
    });

    return entryPoints;
  },
  output: {
    path: path.resolve(__dirname, "../public/dist"),
    filename: "[name].js",
  },
};
