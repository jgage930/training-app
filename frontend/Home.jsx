import React, { useState } from 'react';
import { Modal, Button, ButtonToolbar, Placeholder, Input, DatePicker } from 'rsuite';

function Home() {
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false); 

  return (
    <>
      <ButtonToolbar>
        <Button onClick={handleOpen}> Schedule Workout </Button>
      </ButtonToolbar>

      <Modal open={open} onClose={handleClose}>
        <Modal.Header>
          <Modal.Title> Schedule a Workout </Modal.Title>
        </Modal.Header>

        <Modal.Body>
          <Input placeholder="Name"></Input> 
          <DatePicker></DatePicker>
          <Label>Workout Description</Label>
          <Input as="textarea" rows={10}> </Input>
        </Modal.Body>

        <Modal.Footer>
          <Button onClick={handleClose} appearance="primary">
            Submit
          </Button>
          <Button onClick={handleClose} appearance="subtle">
            Cancel
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  )
}

export default Home;
