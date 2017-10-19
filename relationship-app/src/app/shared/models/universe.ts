export interface Universe {
  hash: string;
  name: string;
}

export namespace Universe {
  export interface Creation {
    name: string;
  }

  export interface Update {
    name: string;
  }
}
