<template lang="pug">
QPage.row.items-center.justify-evenly
  Dialogs
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import Dialogs from 'components/Dialogs.vue';

export default defineComponent({
  name: 'ChatsPage',
  components: {
    QPage,
    Dialogs,
  },
});
</script>
