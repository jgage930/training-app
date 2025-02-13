import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react';

import { Uploader, Button, Message } from 'rsuite';

function ActivityUpload() {
  return (
    <>
      <Message type="success">
        <strong>Success!</strong> Succesfully Uploaded File.
      </Message>

      <Message type="error">
        <strong>Error!</strong> Failed to Upload File.
      </Message>


      <Uploader
        listType="picture-text"
        action="/activity/upload"
        name="fileName"
      >
        <Button>Select files...</Button>
      </Uploader>
    </>
  )
}

export default ActivityUpload;
