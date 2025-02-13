import React, { useState } from 'react'

import 'rsuite/dist/rsuite.min.css';
import { Sidenav, Nav } from 'rsuite';
import CalenderSimpleIcon from '@rsuite/icons/CalenderSimple';
import GroupIcon from '@rsuite/icons/legacy/Group';
import LineChartIcon from '@rsuite/icons/LineChart';
import ArchiveIcon from '@rsuite/icons/Archive';
import StorageIcon from '@rsuite/icons/Storage';

// import WorkoutCalendar from './WorkoutCalendar';
import Home from './Home'
import ActivityUpload from './ActivityUpload';
import AnalyzeWorkout from './AnalyzeWorkout';
import AdminView from './AdminView'

function MainView() {
  // Application Mode
  const [mode, setMode] = useState("calendar")

  function renderContent() {
    switch (mode) {
      case "calendar":
        // return <WorkoutCalendar />
        return <Home />
      case "analyze":
        return <AnalyzeWorkout />
      case "upload":
        return <ActivityUpload />
      case "admin":
        return <AdminView />
    }
  }

  return (
    <div style={{ display: "flex", flexDirection: "row" }}>
      <div style={{ width: 240 }}>
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
              <Nav.Item eventKey="admin" icon={<StorageIcon />}>
                Admin
              </Nav.Item>
            </Nav>
          </Sidenav.Body>
        </Sidenav>
      </div> 
      <div id="content" style={{ width: 1000 }}>{renderContent()}</div>
    </div>
  )
}

export default MainView;
