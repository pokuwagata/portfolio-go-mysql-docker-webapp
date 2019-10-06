import * as React from 'react';

export interface PaginationProps {
  pageNumber: number;
  setPageNumber: (pageNumber: number) => void;
  maxPageNumber: number;
  setMaxPageNumber: (maxPageNumber: number) => void;
  loading: boolean;
}

enum Mode {
  Left,
  Center,
  Right,
  None,
}

enum MoveType {
  Previous,
  Next,
}

const determineMode = (now: number, max: number) => {
  if (max === 1) {
    return Mode.None;
  }
  switch (now) {
    case 1:
      return Mode.Left;
    case max:
      return Mode.Right;
    default:
      return Mode.Center;
  }
};

const getPages = (now: number, max: number) => {
  let pages;
  if (now === 1) {
    // TODO: 畳み込みに書き換え
    pages = [now, now + 1, now + 2];
  } else if (now === max) {
    pages = [now - 2, now - 1, now];
  } else {
    pages = [now - 1, now, now + 1];
  }

  return pages.filter(page => page >= 1 && page <= max);
};

export const Pagination = (props: PaginationProps) => {
  const pages = getPages(props.pageNumber, props.maxPageNumber);
  const mode = determineMode(props.pageNumber, props.maxPageNumber);

  const onClickNumber = (event: React.MouseEvent, number: number) => {
    event.preventDefault();
    props.setPageNumber(number);
  };

  const onClickMove = (event: React.MouseEvent, moveType: MoveType) => {
    event.preventDefault();
    switch (moveType) {
      case MoveType.Next:
        props.setPageNumber(props.pageNumber + 1);
        break;
      case MoveType.Previous:
        props.setPageNumber(props.pageNumber - 1);
        break;
    }
  };

  return (
    <div className="d-flex justify-content-center">
      {props.loading ? null : (
        <nav aria-label="Page navigation example">
          <ul className="pagination">
            {mode !== Mode.Left && mode !== Mode.None && (
              <li className="page-item">
                <a
                  className="page-link"
                  href="#"
                  aria-label="Previous"
                  onClick={(event: React.MouseEvent) =>
                    onClickMove(event, MoveType.Previous)
                  }
                >
                  <span aria-hidden="true">&laquo;</span>
                </a>
              </li>
            )}
            {pages.map(number => (
              <li
                className={
                  'page-item' + (number === props.pageNumber && ' ' + 'active')
                }
                key={number}
              >
                <a
                  className="page-link"
                  href="#"
                  onClick={(event: React.MouseEvent) =>
                    onClickNumber(event, number)
                  }
                >
                  {number}
                </a>
              </li>
            ))}
            {mode !== Mode.Right && mode !== Mode.None && (
              <li className="page-item">
                <a
                  className="page-link"
                  href="#"
                  aria-label="Next"
                  onClick={(event: React.MouseEvent) =>
                    onClickMove(event, MoveType.Next)
                  }
                >
                  <span aria-hidden="true">&raquo;</span>
                </a>
              </li>
            )}
          </ul>
        </nav>
      )}
    </div>
  );
};
