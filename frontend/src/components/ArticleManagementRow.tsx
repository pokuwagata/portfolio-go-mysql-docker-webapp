import * as React from 'react';
import { Link } from 'react-router-dom';

export interface ArticleManagementRowProps {
  id: string;
  title: string;
  updatedAt: string;
  checked: boolean;
  setSelected: (id: string) => void;
}

export const ArticleManagementRow = (props: ArticleManagementRowProps) => {
  const onClickCheck = (
    event: React.ChangeEvent<HTMLInputElement>,
    id: string
  ) => {
    if (event.target.checked) {
      props.setSelected(id);
    } else {
      props.setSelected(null);
    }
  };

  return (
    <div className="d-flex mb-3 align-items-center">
      <div className="mr-1">
        <input
          className="form-check-input position-static"
          type="checkbox"
          id={props.id}
          onChange={event => onClickCheck(event, props.id)}
          checked={props.checked}
        />
      </div>
      <div
        className="mr-5"
        style={{ width: '50%', overflowWrap: 'break-word' }}
      >
        <label
          className="form-check-label"
          htmlFor={props.id}
        >
          {props.title}
        </label>
      </div>
      <div className="">{new Date(props.updatedAt).toLocaleString()}</div>
      <div className="ml-auto">
        <Link to={'/post?id=' + props.id}>
          <button type="button" className="btn btn-primary">
            編集
          </button>
        </Link>
      </div>
    </div>
  );
};
