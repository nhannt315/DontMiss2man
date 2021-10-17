// addDecoratorで参考にした　https://github.com/storybookjs/storybook/issues/5578
// injectFirst https://material-ui.com/guides/interoperability/#controlling-priority-3
// ↓に対応するため、alpha版をいれている
// ref: https://github.com/storybookjs/storybook/issues/12668
import React from 'react';
import { addDecorator } from '@storybook/react';
import { CssBaseline, StylesProvider } from '@material-ui/core';
import 'tailwindcss/tailwind.css';

addDecorator((s) => (
  <StylesProvider injectFirst>
    <CssBaseline />
    {s()}
  </StylesProvider>
));
