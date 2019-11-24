import * as React from 'react';
import * as Const from '../const'

export interface SignoutButtonProps {
  setIsLoggedIn: (state: boolean) => void;
}

export const SignoutButton = (props: SignoutButtonProps) => {
  const onClick = (e: React.MouseEvent) => {
    e.preventDefault();
    localStorage.removeItem(Const.jwtTokenKey)
    props.setIsLoggedIn(false)
  }
  return (
  <li className="nav-item">
    <a className="nav-link" style={{cursor: 'pointer'}} onClick={onClick}>ログアウト</a>
  </li>
)};
