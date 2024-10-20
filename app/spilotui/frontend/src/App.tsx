import { useEffect, useState } from 'react';
import { EventsOn } from '../wailsjs/runtime/runtime.js';
import {
  ConnectServer,
  GetServerStatus,
} from '../wailsjs/go/driver/GrpcDriver.js';
import { StatusResponse } from './app/types/proto/store/status.js';
import { LogResponse } from './app/types/proto/api/v1/log_service.js';
import { apiv1 } from 'wailsjs/go/models.js';
import Pilot from './app/Pilot.js';
import { Theme } from '@radix-ui/themes';
function App() {
  return <Pilot />;
}

export default App;
