import * as React from 'react';
import { Flush } from './Flush';

export enum FlushType {
  SUCCESS,
  ERROR,
  INFO,
}

export type FlushState = {
  isDisplay: boolean;
  type?: FlushType;
  message?: string;
};

export enum FlushActionType {
  VISIBLE,
  HIDDEN,
  FORCE_HIDDEN
}

export type FlushAction = {
  type: FlushActionType;
  payload?: {
    type: FlushType;
    message: string;
  };
};

export type FlushProviderProps = {
  children: React.ReactNode;
};

const reducer = (state: FlushState, action: FlushAction) => {
  switch (action.type) {
    case FlushActionType.VISIBLE:
      switch (action.payload.type) {
        case FlushType.SUCCESS:
          return {
            isDisplay: true,
            type: FlushType.SUCCESS,
            message: action.payload.message,
          };
        case FlushType.ERROR:
          return {
            isDisplay: true,
            type: FlushType.ERROR,
            message: action.payload.message,
          };
      }
    case FlushActionType.HIDDEN:
      switch (state.type) {
        case FlushType.SUCCESS:
          return state;
        case FlushType.ERROR:
          return { isDisplay: false };
        default:
          return state;
      }
    case FlushActionType.FORCE_HIDDEN:
      return { isDisplay: false };
    default:
      return state;
  }
};

export const FlushDispatchContext = React.createContext(null);

export const FlushProvider = (props: FlushProviderProps) => {
  const initState = { isDisplay: false };
  const [state, dispath] = React.useReducer(reducer, initState);

  return (
    <>
      <FlushDispatchContext.Provider value={dispath}>
        <Flush {...state}></Flush>
        {props.children}
      </FlushDispatchContext.Provider>
    </>
  );
};
