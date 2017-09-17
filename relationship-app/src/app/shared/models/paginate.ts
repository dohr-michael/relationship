export interface Paginate<T> {
  items: T[];
  offset: number;
  length: number;
  total: number;
}
