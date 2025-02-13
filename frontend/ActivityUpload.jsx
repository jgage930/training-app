import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react';

import { Uploader, Button } from 'rsuite';

function ActivityUpload() {
  return (
    <Uploader
      listType="picture-text"
      multiple=true
      action="//jsonplaceholder.typicode.com/posts/"
    >
      <Button>Select files...</Button>
    </Uploader>
  t
}

export default ActivityUpload;
