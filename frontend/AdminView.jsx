import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react'

function AdminView() {
  // Call endpoint to get workouts from db.
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

  return (
    
  )
}

export default AdminView;
