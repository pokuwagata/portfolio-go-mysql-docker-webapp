import * as React from 'react';
import { Link } from 'react-router-dom';

export interface ArticleRowProps {
  id: string;
  username: string;
  title: string;
  content: string;
  updatedAt: string;
}

const MAX_CONTENT_LENGTH = 300;

export const ArticleRow = (props: ArticleRowProps) => {
  const isOverMaxLength = (content: string) =>
    content.length > MAX_CONTENT_LENGTH;
  const getViewContent = (content: string) =>
    isOverMaxLength(content)
      ? content.slice(0, MAX_CONTENT_LENGTH) + '...'
      : content;
  return (
    <div className="my-5 pb-4 border-bottom">
      <div className="d-flex align-items-center mb-2">
        <h2>{props.title}</h2>
        <div className="ml-auto">
          {new Date(props.updatedAt).toLocaleString()}
        </div>
      </div>
      <div>
        <p>{getViewContent(props.content)}</p>
        <div className="d-flex">
          {isOverMaxLength(props.content) && (
            <div>
              <Link to={'/article?id=' + props.id}>続きを読む</Link>
            </div>
          )}
          <div className="ml-auto">posted by {props.username}</div>
        </div>
      </div>
    </div>
  );
};
