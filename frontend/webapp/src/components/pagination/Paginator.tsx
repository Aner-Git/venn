import { Container, Pagination } from "react-bootstrap";

interface PaginateParams {
  page?: number;
  pageSize?: number;
}

interface PaginatorParams {
  handleNext: () => void;
  handlePrev: () => void;
  readonly onfirst: boolean;
  readonly onlast: boolean;
  children: JSX.Element;
  size?: "lg" | "sm";
}

/* Pagination helper
 * The first page starts at 1 (not 0)
 */
const Paginator = ({
  children,
  handlePrev,
  handleNext,
  onfirst,
  onlast,
  size = "lg",
}: PaginatorParams) => {
  //if first and last page is disabled no need to paginate
  if (onfirst && onlast) {
    return <>{children}</>;
  }

  return (
    <Container className="p-0">
      {children}
      <Pagination size={size} className="justify-content-center">
        <Pagination.Prev onClick={handlePrev} disabled={onfirst} />
        <Pagination.Next onClick={handleNext} disabled={onlast} />
      </Pagination>
    </Container>
  );
};

Paginator.onPage = (page: number, pageSize: number, totalEntities: number) => {
  return {
    last:
      pageSize >= totalEntities || page === Math.ceil(totalEntities / pageSize),
    first: page === 1,
  };
};

Paginator.defaultPageSize = 20;

export { PaginateParams, Paginator };
