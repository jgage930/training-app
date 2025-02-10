import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react'
import { createRoot } from 'react-dom/client';

import Home from './Home';

export default function Application() {
  return (
    <>
      <div>Application</div>
      <Home />
    </>
  )
}


// Clear the existing HTML content
document.body.innerHTML = '<div id="app"></div>';

// Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application />);
