import * as React from 'react';

export enum FlushType {
  SUCCESS,
  ERROR,
  INFO,
}

export type FlushProps = {
  isDisplay: boolean;
  type: FlushType;
  message: string;
  setFlushState: any;
};

export const Flush = (props: FlushProps) => {
  const convertFlushTypeToClassName = function(type: FlushType) {
    switch (type) {
      case FlushType.SUCCESS:
        return 'alert-success';
      case FlushType.ERROR:
        return 'alert-error';
      case FlushType.INFO:
        return 'alert-info';
      default:
        return '';
    }
  };

  if (props.isDisplay) {
    setTimeout(() => props.setFlushState({ isDisplay: false }), 1500);
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
