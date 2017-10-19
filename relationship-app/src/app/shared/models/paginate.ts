export interface Paginate<T> {
  items: T[];
  page: number;
  size: number;
  total: number;
}
