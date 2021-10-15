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

export type MenuLink = {
  title: string,
  link: string,
}

export interface Chat {
  id: number,
  creator: boolean,
  date: number,
  deactivated: boolean,
  kicked: boolean,
  left: boolean,
  title: string,
}

export type ChatMap = {
  [key: number]: Chat
}

export interface Dialog {
  folderId: number,
  pinned: boolean,
  unreadCount: number
}

export type DialogMap = {
  [key: number]: Dialog
}
