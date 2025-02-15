import 'rsuite/dist/rsuite.min.css';
import React, { useState, useEffect } from 'react'

import { SelectPicker } from 'rsuite'

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
    </>
  )
}

export default AnalyzeWorkout;
