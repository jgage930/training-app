import React, { useState } from 'react';
import { Modal, Button, ButtonToolbar, Placeholder, Input, DatePicker } from 'rsuite';

const dateFormat = 'MM/dd/yyyy';

const today = () => {
  return new Date();
}

const defaultWorkoutName = () => {
  return `Workout ${today().toLocaleDateString("en-US")}`;
}

function Home() {
  const [open, setOpen] = React.useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false); 

  const [workout, setWorkout] = useState({name: defaultWorkoutName(), date: today(), description: ''})

  const handleSubmit = () => {
    console.log("Submitted");

    fetch('/workout', {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(workout)
    })
    .then((data) => console.log("Success"))
    .catch((error) => console.error("Error:", error));
  }

  const handleChange = (key, value) => {
    setWorkout((prevWorkout) => ({ ...prevWorkout, [key]: value}));
  }

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
          <Input 
            placeholder={ defaultWorkoutName() } 
            onChange={(value) => handleChange("name", value)}/> 

          <DatePicker 
            defaultValue={ new Date() } 
            format={ dateFormat }  
            onChange={(value) => handleChange("date", value)}/>

          <Input 
            as="textarea" 
            rows={10} 
            placeholder="Description..." 
            onChange={(value) => handleChange("description", value)}/>
        </Modal.Body>

        <Modal.Footer>
          <Button onClick={handleSubmit} appearance="primary">
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
