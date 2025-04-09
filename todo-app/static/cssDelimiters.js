import styled from 'styled-components';

// JavaScript code injection to dynamically generate CSS
const DynamicButton = styled.button`
  background-color: ${props => (props.primary ? 'blue' : 'gray')};
  color: ${props => (props.disabled ? 'lightgray' : 'white')};
  font-size: ${props => `${props.size || 16}px`};
  margin: 10px;
  padding: ${props => props.large ? '12px 24px' : '8px 16px'};

  &:hover {
    background-color: ${props => (props.primary ? 'darkblue' : 'darkgray')};
  }
`;

// Usage
<DynamicButton primary size={20}>Primary Button</DynamicButton>
<DynamicButton disabled>Disabled Button</DynamicButton>

const Container = styled.div`
  padding: ${props => props.padding || '10px'};
  color: ${props => (props.isActive ? 'blue' : 'gray')};
`;