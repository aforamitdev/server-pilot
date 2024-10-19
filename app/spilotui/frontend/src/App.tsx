import { useEffect, useState } from 'react';
import { EventsOn } from '../wailsjs/runtime/runtime.js';
import {
  ConnectServer,
  GetServerStatus,
} from '../wailsjs/go/driver/GrpcDriver.js';
import { StatusResponse } from './app/types/proto/store/status.js';
import { LogResponse } from './app/types/proto/api/v1/log_service.js';
import { apiv1 } from 'wailsjs/go/models.js';
function App() {
  const [state, setState] = useState<apiv1.GetStatusResponse | null>(null);
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
        GetServerStatus().then((res: apiv1.GetStatusResponse) => {
          console.log(res, 'resss');
          setState(res);
        });
      })
      .catch((err) => {
        console.log(err);
      });
  }, []);

  return (
    <div id='App' className='bg-white h-screen w-full'>
      {JSON.stringify(state?.system)}
    </div>
  );
}

export default App;
