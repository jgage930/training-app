import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react';

import { Uploader, Button, Message } from 'rsuite';

function ActivityUpload() {
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const handleSuccess = () => {
    setSuccess(true);
    setTimeout(() => setSuccess(false), 3000);
  }

  const handleError = () => {
    setError(true);
    setTimeout(() => setError(false), 3000);
  }

  return (
    <>
      {success && (
        <Message type="success">
          <strong>Success!</strong> Succesfully Uploaded File.
        </Message>
      )}

      {error && (
        <Message type="error">
          <strong>Error!</strong> Failed to Upload File.
        </Message>
      )}

      <Uploader
        multiple={true}
        listType="picture-text"
        action="/activity/upload"
        name="fileName"
        onSuccess={(r, f) => handleSuccess()}
        onError={(r, f) => handleError()}
      >
        <Button>Select files...</Button>
      </Uploader>
    </>
  )
}

export default ActivityUpload;
