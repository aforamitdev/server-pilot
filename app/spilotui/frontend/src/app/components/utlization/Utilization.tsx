import React from 'react';
import {
  Badge,
  Box,
  Container,
  Flex,
  IconButton,
  Text,
} from '@radix-ui/themes';
import {
  ArrowLeft,
  ChevronLeft,
  MoveLeft,
  CheckCircle2,
  ChevronDown,
} from 'lucide-react';

type Props = {};

const Utilization = (props: Props) => {
  return (
    <Box className='h-screen bg-gray-100'>
      <Container className='py-5 mx-5'>
        <div className='flex  items-center'>
          <ArrowLeft size={16} absoluteStrokeWidth />
          <Text className='px-2 text-sm '>Dashboard</Text>
        </div>

        {/* status */}

        <div className='flex items-center space-x-2 pt-4'>
          <div>
            <Badge className='py-2 rounded-full' color='green'>
              <CheckCircle2 color='green' />
            </Badge>
          </div>
          {/* status */}
          <div>
            <div className='flex items-center'>
              <p className='text-xl font-medium leading-5 text-sm'>127.0.0.1</p>
              <ChevronDown size={16} />
            </div>
            <p className='text-sm text-green-700'>check every 30 seconds</p>
          </div>
          {/* details */}
        </div>

        <div className='flex py-4  justify-between'>
          <div>
            <div className='font-semibold'>IP Details </div>
          </div>
          <div>
            <div className='font-semibold'>Driver Details</div>
          </div>
        </div>

        <div className='bg-white rounded-md h-64 px-2 py-2'>
          <div>
            <div className='font-medium'>System Utilization</div>
            <div className='text-xs text-gray-600'>
              Data is updated every 30 secons{' '}
            </div>
          </div>
        </div>
      </Container>
    </Box>
  );
};

export default Utilization;
