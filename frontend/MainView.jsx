import React, { useState } from 'react'

import 'rsuite/dist/rsuite.min.css';
import { Sidenav, Nav } from 'rsuite';
import CalenderSimpleIcon from '@rsuite/icons/CalenderSimple';
import GroupIcon from '@rsuite/icons/legacy/Group';
import LineChartIcon from '@rsuite/icons/LineChart';
import ArchiveIcon from '@rsuite/icons/Archive';

import WorkoutCalendar from './WorkoutCalendar';
import ActivityUpload from './ActivityUpload';
import AnalyzeWorkout from './AnalyzeWorkout';



function MainView() {
  // Application Mode
  const [mode, setMode] = useState("calendar")

  function renderContent() {
    switch (mode) {
      case "calendar":
        return <WorkoutCalendar />
      case "analyze":
        return <AnalyzeWorkout />
      case "upload":
        return <ActivityUpload />
    }
  }

  return (
    <div style={{ width: 240, display: "flex", flexDirection: "column" }}>
      <Sidenav>
        <Sidenav.Body>
          <Nav defaultActiveKey="calendar" onSelect={(eventKey => setMode(eventKey))}>
            <Nav.Item eventKey="calendar" icon={<CalenderSimpleIcon />}>
              Calendar
            </Nav.Item>
            <Nav.Item eventKey="analyze" icon={<LineChartIcon />}>
              Analyze Workout
            </Nav.Item>
            <Nav.Item eventKey="upload" icon={<ArchiveIcon />}>
              Upload Activities
            </Nav.Item>
          </Nav>
        </Sidenav.Body>
      </Sidenav>
      <div id="content">{renderContent()}</div>
    </div> 
  )
}

export default MainView;
