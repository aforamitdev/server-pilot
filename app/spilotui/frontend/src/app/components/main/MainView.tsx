import React from 'react';
import {
  ResizableHandle,
  ResizablePanel,
  ResizablePanelGroup,
} from './ResizableContainer';
import { cn } from '@/lib/utils';
import Utilization from '../utlization/Utilization';

type Props = {
  defaultLayout?: number[];
  defaultCollapsed?: boolean;
  navCollapsedSize?: number;
};

function MainView({
  defaultLayout = [20, 32, 48],
  defaultCollapsed = false,
  navCollapsedSize = 30,
}: Props) {
  const [isCollapsed, setIsCollapsed] = React.useState(defaultCollapsed);

  return (
    <ResizablePanelGroup
      direction='horizontal'
      onLayout={(sizes: number[]) => {
        document.cookie = `react-resizable-panels:layout:mail=${JSON.stringify(
          sizes
        )}`;
      }}
      className='h-full  items-stretch'
    >
      <ResizablePanel
        defaultSize={defaultLayout[0]}
        collapsedSize={navCollapsedSize}
        collapsible={true}
        minSize={15}
        maxSize={20}
        onCollapse={() => {
          setIsCollapsed(true);
          document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
            true
          )}`;
        }}
        onResize={() => {
          setIsCollapsed(false);
          document.cookie = `react-resizable-panels:collapsed=${JSON.stringify(
            false
          )}`;
        }}
        className={cn(
          isCollapsed && 'min-w-[50px] transition-all duration-300 ease-in-out'
        )}
      >
        <div
          className={cn(
            'flex h-[52px] items-center justify-center',
            isCollapsed ? 'h-[52px]' : 'px-2'
          )}
        >
          asas
        </div>
      </ResizablePanel>
      <ResizableHandle withHandle />
      <ResizablePanel defaultSize={defaultLayout[1]} minSize={30}>
        <Utilization />
      </ResizablePanel>
    </ResizablePanelGroup>
  );
}

export default MainView;
