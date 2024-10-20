import React from 'react';
import { Box, Container, Flex, IconButton, Text } from '@radix-ui/themes';
import { ChevronLeft, MoveLeft } from 'lucide-react';
import { MagnifyingGlassIcon } from '@radix-ui/react-icons';

type Props = {};

const Utilization = (props: Props) => {
  return (
    <Box
      style={{ background: 'var(--gray-a2)', borderRadius: 'var(--radius-3)' }}
      className='h-screen'
    >
      <Container className='py-5 mx-5'>
        <div className='flex text-gray-300'>
          <MoveLeft size={20} />
          <div className='px-2 text-red-600 font-bold'>Dashboard asas</div>
        </div>
      </Container>
    </Box>
  );
};

export default Utilization;
