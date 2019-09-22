import * as React from 'react';
import { ArticleManagementRow } from './ArticleManagementRow';

export interface ArticleManagementProps {}

export const ArticleManagement = (props: ArticleManagementProps) => {
  const articleList = [
    {
      id: '1',
      title: '記事のタイトル',
      updatedAt: '2019/09/19/20:00:00',
    },
    {
      id: '2',
      title: '記事のタイトル',
      updatedAt: '2019/09/19/20:00:00',
    },
    {
      id: '3',
      title: '記事のタイトル',
      updatedAt: '2019/09/19/20:00:00',
    },
  ];

  return (
    <div>
      <h1 className="mb-3">記事の管理</h1>
      <div className="container mb-5">
        <div className="row">
          <div className="ml-md-auto">
            <button type="button" className="btn btn-danger">
              選択した記事を削除
            </button>
          </div>
        </div>
      </div>
      <div className="container">
        {articleList.map(article => (
          <ArticleManagementRow
            id={article.id}
            title={article.title}
            updatedAt={article.updatedAt}
          />
        ))}
      </div>
    </div>
  );
};
