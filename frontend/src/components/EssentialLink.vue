<template lang="pug">
QItem(
  clickable
  tag="a"
  :to="{name: link}"
)
  QItemLabel {{ title }}
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QItem, QItemLabel } from "quasar";

export default defineComponent({
  name: 'EssentialLink',
  components: {
    QItem,
    QItemLabel,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
    link: {
      type: String,
      required: true,
    },
  },
});
</script>
