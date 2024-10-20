import React from 'react';
import { createRoot } from 'react-dom/client';
import './output.css';
import '@radix-ui/themes/styles.css';
import App from './App';
import { Theme, ThemePanel } from '@radix-ui/themes';

const container = document.getElementById('root');

const root = createRoot(container!);

root.render(
  <React.StrictMode>
    <div id='App' className=''>
      <Theme
        accentColor='crimson'
        grayColor='sand'
        radius='large'
        scaling='95%'
      >
        <App />
        {/* <ThemePanel /> */}
      </Theme>
    </div>
  </React.StrictMode>
);
