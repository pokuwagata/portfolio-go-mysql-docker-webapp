import * as React from 'react';
import { NavLink } from 'react-router-dom';

export interface HeaderButtonProps {
  name: string;
  path: string;
}

export const HeaderButton = (props: HeaderButtonProps) => (
  <li className="nav-item">
    <NavLink
      exact
      to={props.path}
      activeClassName="selected"
      className="nav-link"
    >
      {props.name}
    </NavLink>
  </li>
);
