import React from 'react';
import { Badge } from '@radix-ui/themes';
import MainView from './components/main/MainView';
import { TooltipProvider } from '@radix-ui/react-tooltip';

type Props = {};

const Pilot = (props: Props) => {
  return (
    <TooltipProvider delayDuration={0}>
      <div className=' w-full h-screen  font-Inter'>
        <div className='border-b  '>
          <div className='flex  items-center px-4 py-2 space-x-2 '>
            <Badge>linux | ubuntu 20.10</Badge> <Badge>127.0.0.0.1</Badge>
          </div>
          <div className='h-screen border-t'>
            <MainView />
          </div>
        </div>
      </div>
    </TooltipProvider>
  );
};

export default Pilot;
