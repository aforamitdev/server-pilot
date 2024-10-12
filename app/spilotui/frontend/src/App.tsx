import { useEffect, useState } from 'react';
import { EventsOn } from '../wailsjs/runtime/runtime.js';

function App() {
  const [state, setState] = useState('');
  useEffect(() => {
    EventsOn('LOG:RECEIVED', (payload) => {
      console.log(payload);
      setState(payload);
    });
  }, []);
  return (
    <div id='App' className='bg-white h-screen w-full'>
      aasasas {state}
    </div>
  );
}

export default App;
