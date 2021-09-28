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
      QToolbarTitle Telegram client managment system
  QDrawer.bg-grey-1(
    v-model="leftDrawerOpen"
    show-if-above
    bordered
  )
    QList
      QItemLabel.text-grey-8(
        header
      ) Essential Links
      EssentialLink(
        v-for="link in essentialLinks"
        :key="link.title"
        v-bind="link"
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

const linksList = [
  {
    title: 'Main',
    caption: 'main',
    link: '#',
  },
];

import { defineComponent, ref } from 'vue';

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
      essentialLinks: linksList,
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
    };
  },
});
</script>
