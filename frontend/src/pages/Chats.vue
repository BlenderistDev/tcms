<template lang="pug">
QPage.row.items-center.justify-evenly
  Chats
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import Chats from 'components/Chats.vue';

export default defineComponent({
  name: 'ChatsPage',
  components: {
    QPage,
    Chats,
  },
});
</script>
