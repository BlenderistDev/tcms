export interface Meta {
  totalCount: number;
}

export interface User {
  id: number,
  phone: string,
  name: string
}

export interface Contact {
  id: number
  username: string,
  phone: string,
  firstName: string,
  lastName: string,
  mutualContact: boolean,
  bot: boolean,
  deleted: boolean,
  accessHash: number,
}

export type ContactMap = {
  [key: number]: Contact
}
