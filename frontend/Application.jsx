import React, { useState } from 'react'
import { createRoot } from 'react-dom/client';
import 'rsuite/dist/rsuite.min.css';




export default function Application() {
  return (
    <>
      <div>Application</div>
    </>
  )
}


// Clear the existing HTML content
document.body.innerHTML = '<div id="app"></div>';

// Render your React component instead
const root = createRoot(document.getElementById('app'));
root.render(<Application />);
