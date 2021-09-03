<template lang="pug">
QPage.row.items-center.justify-evenly
  Contacts
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import Contacts from 'components/Contacts.vue';

export default defineComponent({
  name: 'PageLogin',
  components: {
    QPage,
    Contacts,
  },
});
</script>
