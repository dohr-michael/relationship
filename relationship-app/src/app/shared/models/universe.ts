export interface Universe {
  id: string;
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
