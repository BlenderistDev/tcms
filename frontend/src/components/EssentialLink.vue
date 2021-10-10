<template lang="pug">
QItem(
  clickable
  tag="a"
  :to="{name: link.link}"
)
  QItemLabel {{ link.title }}
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { QItem, QItemLabel } from 'quasar';
import { MenuLink } from 'components/models';

export default defineComponent({
  name: 'EssentialLink',
  components: {
    QItem,
    QItemLabel,
  },
  props: {
    link: {
      type: Object as PropType<MenuLink>,
      required: true,
    },
  },
});
</script>
