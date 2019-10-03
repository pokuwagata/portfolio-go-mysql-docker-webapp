import * as React from 'react';
import { ArticleManagementRow } from './ArticleManagementRow';
import { FlushState } from './App';
import { FlushType } from './Flush';
import { Pagination } from './Pagination';

export interface ArticleManagementProps {
  isLoggedIn: boolean;
  setIsLoggedIn: (state: boolean) => void;
  setFlushState: (state: FlushState) => void;
}

export const ArticleManagement = (props: ArticleManagementProps) => {
  const [loading, setLoading] = React.useState(true);
  const [articleList, setArticleList] = React.useState([]);
  const [pageNumber, setPageNumber] = React.useState(1);
  const [maxPageNumber, setMaxNumber] = React.useState(10);

  React.useEffect(() => {
    fetch('api/admin/articles?number=' + pageNumber, {
      method: 'GET',
      headers: {
        'content-type': 'application/json',
        Authorization: 'Bearer ' + localStorage.getItem('portfolio-jwt-token'),
      },
    })
      .then(res => {
        return new Promise(resolve =>
          res.json().then(json => {
            resolve({
              ok: res.ok,
              json,
            });
            setLoading(false);
          })
        );
      })
      .then(res => {
        // TODO: as any以外の方法
        if ((res as any).ok) {
          props.setFlushState({ isDisplay: false });
          setArticleList((res as any).json.articles);
          (res as any).json.maxNumber &&
            setMaxNumber((res as any).json.maxNumber);
        } else {
          throw new Error((res as any).json.message);
        }
      })
      .catch(error => {
        props.setFlushState({
          isDisplay: true,
          type: FlushType.ERROR,
          message: '記事の取得に失敗しました。' + error,
        });
        setArticleList([]);
      });
  }, [pageNumber]);

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
      <div className="container mb-5" style={{ minHeight: '35vh' }}>
        {loading ? (
          <p>loading...</p>
        ) : articleList.length === 0 ? (
          <p>記事が見つかりません</p>
        ) : (
          articleList.map(article => (
            <ArticleManagementRow
              key={article.id}
              id={article.id}
              title={article.title}
              updatedAt={article.updatedAt}
            />
          ))
        )}
      </div>
      <Pagination
        pageNumber={pageNumber}
        setPageNumber={setPageNumber}
        maxPageNumber={maxPageNumber}
        setMaxPageNumber={setPageNumber}
      />
    </div>
  );
};
