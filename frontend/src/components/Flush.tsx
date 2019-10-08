import * as React from 'react';
import { FlushDispatchContext, FlushActionType } from './FlushProvider';

export enum FlushType {
  SUCCESS,
  ERROR,
  INFO,
}

export type FlushProps = {
  isDisplay: boolean;
  type?: FlushType;
  message?: string;
};

const FLUSH_HIDDEN_TIME = 3000;

export const Flush = (props: FlushProps) => {
  const flushDispatch = React.useContext(FlushDispatchContext);

  const convertFlushTypeToClassName = function(type: FlushType) {
    switch (type) {
      case FlushType.SUCCESS:
        return 'alert-success';
      case FlushType.ERROR:
        return 'alert-danger';
      case FlushType.INFO:
        return 'alert-info';
      default:
        return '';
    }
  };

  if (props.isDisplay && props.type === FlushType.SUCCESS) {
    setTimeout(
      () => flushDispatch({ type: FlushActionType.FORCE_HIDDEN }),
      FLUSH_HIDDEN_TIME
    );
  }

  return (
    props.isDisplay && (
      <div
        className={'alert ' + convertFlushTypeToClassName(props.type)}
        role="alert"
      >
        {props.message}
      </div>
    )
  );
};
