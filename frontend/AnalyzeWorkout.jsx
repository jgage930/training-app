import 'rsuite/dist/rsuite.min.css';
import React, { useState, useEffect } from 'react'

import { SelectPicker, Panel, Stack, Stat } from 'rsuite'

function ActivityStats({id}) {
  const [stats, setStats] = useState({});

  const fetchStats = () => {
    fetch(`/stats/activity/${id}`, { method: "GET" })
    .then((response) => {
      return response.json();
    })
    .then((stats) => {
        setStats(stats);
    });
  }

  useEffect(() => { fetchStats() }, [])

  return (
    <Panel header="Activity Stats" bordered>
      <Stack wrap spacing={20}>
        <Stat bordered style={{ width: 200 }} >
          <Stat.Label>Total Distance</Stat.Label>
          <Stat.Value 
            value={stats.total_distance}
            formatOptions={{
              style: "unit",
              unit: 'mile'
            }}
          />
        </Stat> 
        
        <Stat bordered style={{ width: 200 }}>
          <Stat.Label>Average Speed</Stat.Label>
          <Stat.Value 
            value={stats.average_speed}
            formatOptions={{
              style: "unit",
              unit: 'mile-per-hour'
            }}
          />
        </Stat> 

        <Stat bordered style={{ width: 200 }}>
          <Stat.Label>Max Speed</Stat.Label>
          <Stat.Value 
            value={stats.max_speed}
            formatOptions={{
              style: "unit",
              unit: 'mile-per-hour'
            }}
          />
        </Stat> 

        <Stat bordered style={{ width: 200 }}>
          <Stat.Label>Average Heart Rate</Stat.Label>
          <Stat.Value 
            value={stats.average_heart_rate}
          />
        </Stat> 
      

      </Stack>
    </Panel>
  )
}

function AnalyzeWorkout() {
  const [activities, setActivities] = useState([]);
  
  const fetchActivities = () => {
    fetch('/activity', { method: "GET" })
    .then((response) => {
      return response.json();
    })
    .then((activities) => {
        setActivities(activities);
    });
  }

  useEffect(() => {fetchActivities()}, []);

  const [selectOptions, setSelectOptions] = useState([]);
  useEffect(() => {
    const options = activities.map((item) => ({"label": item.file_path, "value": item.id}));
    setSelectOptions(options);
  }, [activities]);

  // Activity Id to diplay analysis for.
  const [activityId, setActivityId] = useState(null);

  const handleSelect = (value, item) => {
    setActivityId(value); 
  }

  useEffect(() => {console.log(activityId)});

  return (
    <>
      <SelectPicker 
        data={selectOptions} 
        onSelect={(value, item) => {handleSelect(value, item)}}
      />
      {activityId && (
        <ActivityStats id={activityId}/>
      )}
    </>
  )
}

export default AnalyzeWorkout;
