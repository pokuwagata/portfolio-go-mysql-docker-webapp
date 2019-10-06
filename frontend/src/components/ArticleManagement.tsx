import * as React from 'react';
import { ArticleManagementRow } from './ArticleManagementRow';
import { FlushState } from './App';
import { FlushType } from './Flush';
import { Pagination } from './Pagination';
import { ArticleRemoveButton } from './ArticleRemoveButton';

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
  const [selected, setSelected] = React.useState(null);

  const fetchArticles = async (pageNumber: number) => {
    try {
      const res = await fetch('api/admin/articles?number=' + pageNumber, {
        method: 'GET',
        headers: {
          'content-type': 'application/json',
          Authorization:
            'Bearer ' + localStorage.getItem('portfolio-jwt-token'),
        },
      });
      const json = await res.json();
      if (res.ok) {
        props.setFlushState({ isDisplay: false });
        setArticleList(json.articles);
        json.maxNumber && setMaxNumber(json.maxNumber);
        setLoading(false);
      } else {
        throw new Error(json.message);
      }
    } catch (error) {
      props.setFlushState({
        isDisplay: true,
        type: FlushType.ERROR,
        message: '記事の取得に失敗しました。' + error,
      });
      setArticleList([]);
      setLoading(false);
    }
  };

  React.useEffect(() => {
    fetchArticles(pageNumber);
  }, [pageNumber]);

  return (
    <div>
      <h1 className="mb-3">記事の管理</h1>
      <div className="container mb-5">
        <div className="row">
          <div className="ml-md-auto">
            <ArticleRemoveButton
              setFlushState={props.setFlushState}
              fetchArticles={fetchArticles}
              id={selected}
            />
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
              checked={article.id === selected}
              setSelected={setSelected}
            />
          ))
        )}
      </div>
      <Pagination
        pageNumber={pageNumber}
        setPageNumber={setPageNumber}
        maxPageNumber={maxPageNumber}
        setMaxPageNumber={setPageNumber}
        loading={loading}
      />
    </div>
  );
};
