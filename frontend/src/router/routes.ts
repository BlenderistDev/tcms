import { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        path: '',
        component: () => import('pages/Index.vue'),
      },
      {
        path: 'login',
        component: () => import('pages/Login.vue'),
      },
      {
        path: 'contacts',
        name: 'contacts',
        component: () => import('pages/Contacts.vue'),
      },
      {
        path: 'chats',
        name: 'chats',
        component: () => import('pages/Chats.vue'),
      },
      {
        path: 'dialogs',
        name: 'dialogs',
        component: () => import('pages/Dialogs.vue'),
      },
      {
        path: 'send',
        name: 'send',
        component: () => import('pages/Send.vue'),
      },
    ],
  },
  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue'),
  },
];

export default routes;
