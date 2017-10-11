export namespace neo {
  export interface Node<T> {
    NodeIdentity: number;
    Labels: string[];
    Properties: T;
  }
}
