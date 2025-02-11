import React, { useState, useEffect } from 'react';
import { Modal, Button, ButtonToolbar, Placeholder, Input, DatePicker, Calendar, List } from 'rsuite';

const dateFormat = 'MM/dd/yyyy';

const today = () => {
  return new Date();
}

const isSameDay = (day1, day2) => {
  // Check if the day month and year are equivalent.
  const day = day1.getDate() === day2.getDate();
  const month = day1.getMonth() === day2.getMonth();
  const year = day1.getYear() === day2.getYear();

  return day && month && year
}

const defaultWorkoutName = () => {
  const dateStr = today().toLocaleDateString("en-US");
  return `Workout ${dateStr}`;
}

function CalendarCell(date, workouts) {
  const todaysWorkouts = workouts.filter((w) => isSameDay(w.date, date));
  
  return (
     <List>
       {todaysWorkouts.map(workout => (
         <List.Item key={workout.id}>
           <strong>{workout.name}</strong>
         </List.Item>
      ))}
    </List>
  )
}

function WorkoutCalendar() {
  const [workouts, setWorkouts] = useState([]);
  
  useEffect(() => {
    fetch('/workout', {
      method: "GET" 
    })
    .then((response) => {
      return response.json();
    })
    .then((workouts) => {
        const parsedObjects = workouts.map(w => ({...w, date: new Date(w.date)}));
        setWorkouts(parsedObjects)
    });
  }, []);

  return <Calendar bordered renderCell={(date) => CalendarCell(date, workouts)} />
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

      <WorkoutCalendar />
    </>
  )
}

export default Home;
