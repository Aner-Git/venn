const CracoLessPlugin = require("craco-less");

module.exports = {
  plugins: [
    {
      plugin: CracoLessPlugin,
      options: {
        lessLoaderOptions: {
          lessOptions: {
            modifyVars: {
              "@layout-header-background": "#32835C",
              "@menu-bg": "#32835C",
              "@primary-color": "#32835C",
              "@menu-item-color": "white",
              "@menu-item-group-height": "20px",
              "@menu-submenu-bg": "#32835C",
              "@menu-dark-bg": "white",
              "@menu-dark-item-color": "#262626",
              "@menu-highlight-color": "#dedede",
            },
            javascriptEnabled: true,
          },
        },
      },
    },
  ],
};
