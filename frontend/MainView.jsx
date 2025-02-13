import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react'
import { Sidenav, Nav } from 'rsuite';
import CalenderSimpleIcon from '@rsuite/icons/CalenderSimple';
import GroupIcon from '@rsuite/icons/legacy/Group';
import LineChartIcon from '@rsuite/icons/LineChart';
import ArchiveIcon from '@rsuite/icons/Archive';


function MainView() {
  // 1: Calendar
  // 2: Analyze Workout
  // 3: Activity Upload
  const [mode, setMode] = useState("1")

  return (
    <div style={{ width: 240, display: "flex", flexDirection: "column" }}>
      <Sidenav>
        <Sidenav.Body>
          <Nav defaultActiveKey="calendar">
            <Nav.Item eventKey="calendar" icon={<CalenderSimpleIcon />}>
              Calendar
            </Nav.Item>
            <Nav.Item eventKey="analyze" icon={<LineChartIcon />}>
              Analyze Workout
            </Nav.Item>
            <Nav.Item eventKey="activity_upload" icon={<ArchiveIcon />}>
              Upload Activities
            </Nav.Item>
          </Nav>
        </Sidenav.Body>
      </Sidenav>
    </div> 
  )
}

export default MainView;
