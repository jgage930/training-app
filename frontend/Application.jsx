import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react'
import { createRoot } from 'react-dom/client';

import Home from './Home';
import MainView from './MainView'

export default function Application() {
  return (
    <>
      <MainView />
    </>
  )
}


// Clear the existing HTML content
document.body.innerHTML = '<div id="app"></div>';

// Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application />);
