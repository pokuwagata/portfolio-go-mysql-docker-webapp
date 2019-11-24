import * as React from 'react';
import { Redirect, Route } from 'react-router-dom';

export interface PrivateRouteProps {
  isLoggedIn: boolean;
  path: string;
  children: React.ReactNode;
}

export const PrivateRoute = (props: PrivateRouteProps) => {
  return (
    <Route
      path={props.path}
      render={() =>
        props.isLoggedIn ? props.children : <Redirect to="/"></Redirect>
      }
    ></Route>
  );
};
