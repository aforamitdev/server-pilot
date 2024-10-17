import { useEffect, useState } from 'react';
import { EventsOn } from '../wailsjs/runtime/runtime.js';
import {
  ConnectServer,
  GetServerStatus,
} from '../wailsjs/go/driver/GrpcDriver.js';
function App() {
  const [state, setState] = useState('');
  useEffect(() => {
    EventsOn('LOG:RECEIVED', (payload) => {
      console.log(payload);
      setState(payload);
    });

    // ConnectServer('127.0.0.1', '50051')
    //   .then((res) => {
    //     console.log(res);
    //     console.log('res');
    //   })
    //   .catch((err) => {
    //     console.log(err);
    //     console.log('err');
    //   });

    ConnectServer('127.0.0.1', '50051')
      .then((res) => {
        console.log(res);
        GetServerStatus().then((res) => {
          console.log(res, 'resss');
        });
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <div id='App' className='bg-white h-screen w-full'>
      aasasas {state}
    </div>
  );
}

export default App;
