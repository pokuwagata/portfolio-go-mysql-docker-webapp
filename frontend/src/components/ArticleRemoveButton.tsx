import * as React from 'react';
import { FlushType } from './Flush';
import { FlushDispatchContext, FlushActionType } from './FlushProvider';
import * as Const from '../const'

export interface ArticleRemoveButtonProps {
  fetchArticles: (pageNumber: number) => void;
  id: number;
}

export const ArticleRemoveButton = (props: ArticleRemoveButtonProps) => {
  const flushDispath = React.useContext(FlushDispatchContext);

  const isActivated = !!props.id;

  const onClickButton = async (event: React.MouseEvent, id: number) => {
    event.preventDefault();
    if(!isActivated) {
      return
    }

    try {
      const res = await fetch('api/admin/article?id=' + id, {
        method: 'DELETE',
        headers: {
          Authorization:
            'Bearer ' + localStorage.getItem(Const.jwtTokenKey),
        },
      });
      const json = await res.json();
      if (res.ok) {
        flushDispath({
          type: FlushActionType.VISIBLE,
          payload: {
            type: FlushType.SUCCESS,
            message: '記事の削除に成功しました。',
          },
        });
        props.fetchArticles(1);
      } else {
        throw new Error(json.message);
      }
    } catch (error) {
      flushDispath({
        type: FlushActionType.VISIBLE,
        payload: {
          type: FlushType.ERROR,
          message: '記事の削除に失敗しました。' + error,
        },
      });
    }
  };

  return (
    <button
      type="button"
      className={`btn btn-danger ${isActivated ? "" : "disabled"}`}
      style={{cursor: isActivated ? "pointer" : "not-allowed"}}
      onClick={(event: React.MouseEvent) => onClickButton(event, props.id)}
    >
      選択した記事を削除
    </button>
  );
};
