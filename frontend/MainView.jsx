import 'rsuite/dist/rsuite.min.css';
import React, { useState } from 'react'
import { Sidenav, Nav } from 'rsuite';
import CalenderSimpleIcon from '@rsuite/icons/CalenderSimple';
import GroupIcon from '@rsuite/icons/legacy/Group';
import LineChartIcon from '@rsuite/icons/LineChart';
import ArchiveIcon from '@rsuite/icons/Archive';


function MainView() {
  return (
    <div style={{ width: 240, height: '100vh' }}>
      <Sidenav>
        <Sidenav.Body>
          <Nav activeKey="1">
            <Nav.Item eventKey="1" icon={<CalenderSimpleIcon />}>
              Calendar
            </Nav.Item>
            <Nav.Item eventKey="2" icon={<LineChartIcon />}>
              Analyze Workout
            </Nav.Item>
            <Nav.Item eventKey="3" icon={<ArchiveIcon />}>
              Upload Activities
            </Nav.Item>
          </Nav>
        </Sidenav.Body>
      </Sidenav>
    </div> 
  )
}

export default MainView;
