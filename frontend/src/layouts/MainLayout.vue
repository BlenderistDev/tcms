<template lang="pug">
QLayout(view="lHh Lpr lFf")
  QHeader(elevated)
    QToolbar
      QBtn(
        flat
        dense
        round
        icon="menu"
        aria-label="Menu"
        @click="toggleLeftDrawer"
      )
      QToolbarTitle Telegram client management system
  QDrawer.bg-grey-1(
    v-model="leftDrawerOpen"
    show-if-above
    bordered
  )
    QList
      QItemLabel.text-grey-8(
        header
      ) Menu
      EssentialLink(
        v-for="link in links"
        :key="link.title"
        :link="link"
      )
  QPageContainer
    router-view
</template>

<script lang="ts">
import EssentialLink from 'components/EssentialLink.vue';
import {
  QBtn,
  QToolbarTitle,
  QToolbar,
  QHeader,
  QList,
  QItemLabel,
  QDrawer,
  QPageContainer,
  QLayout
} from 'quasar';

import { defineComponent, ref } from 'vue';
import { MenuLink } from 'components/models';

const links: MenuLink[] = [
  {
    title: 'Main',
    link: '',
  },
  {
    title: 'Contacts',
    link: 'contacts',
  },
  {
    title: 'Chats',
    link: 'chats',
  },
  {
    title: 'Dialogs',
    link: 'dialogs',
  },
  {
    title: 'Send',
    link: 'send',
  },
];

export default defineComponent({
  name: 'MainLayout',
  components: {
    EssentialLink,
    QBtn,
    QToolbarTitle,
    QToolbar,
    QHeader,
    QList,
    QItemLabel,
    QDrawer,
    QPageContainer,
    QLayout,
  },

  setup() {
    const leftDrawerOpen = ref(false);

    return {
      links,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
    };
  },
});
</script>
