import * as React from 'react';
import { Link } from 'react-router-dom';

export interface HeaderButtonProps {
  name: string;
  path: string;
}

export const HeaderButton = (props: HeaderButtonProps) => (
  <Link to={props.path}>{props.name}</Link>
);
