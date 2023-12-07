/** @type { import('@storybook/server').Preview } */
const preview = {
  parameters: {
    server: {
      url: `http://localhost:1323/storybook_preview`,
    },
    actions: { argTypesRegex: "^on[A-Z].*" },
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
};

export default preview;
