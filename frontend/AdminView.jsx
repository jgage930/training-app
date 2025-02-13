import 'rsuite/dist/rsuite.min.css';
import React, { useState, useEffect } from 'react';
import { Table, Button } from 'rsuite';

const { Column, HeaderCell, Cell } = Table;

// Sample JSON data
const workoutData = [
  {
    id: 1,
    name: "Test workout",
    date: "2025-02-10T18:01:58.665Z",
    description: "testing",
  },
  {
    id: 4,
    name: "Test workout",
    date: "2025-02-10T18:01:58.665Z",
    description: "testing",
  },
];

function AdminView() {
  // Call endpoint to get workouts from db.
  const [workouts, setWorkouts] = useState(workoutData);
  
  useEffect(() => {
    fetch('/workout', {
      method: "GET" 
    })
    .then((response) => {
      return response.json();
    })
    .then((workouts) => {
        //const parsedObjects = workouts.map(w => ({...w, date: new Date(w.date)}));
        setWorkouts(workouts)
    });
  }, []);

  return (
    <Table
      height={400}
      data={workouts}
      bordered
      cellBordered
      autoHeight
      hover
    >
      <Column width={50} align="center" fixed>
        <HeaderCell>ID</HeaderCell>
        <Cell dataKey="id" />
      </Column>

      <Column flexGrow={1}>
        <HeaderCell>Name</HeaderCell>
        <Cell dataKey="name" />
      </Column>

      <Column flexGrow={1}>
        <HeaderCell>Date</HeaderCell>
        <Cell dataKey="date">
          {(rowData) => new Date(rowData.date).toLocaleString()}
        </Cell>
      </Column>

      <Column flexGrow={1}>
        <HeaderCell>Description</HeaderCell>
        <Cell dataKey="description" />
      </Column>
      <Column width={80} fixed="right">
        <HeaderCell>Delete</HeaderCell>

        <Cell style={{ padding: '6px' }}>
          {rowData => (
            <Button onClick={() => alert(`id:${rowData.id}`)} color='red'>
              Delete
            </Button>
          )}
        </Cell>
      </Column>
    </Table>
  )
}

export default AdminView;
