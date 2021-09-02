<template lang="pug">
QPage.row.items-center.justify-evenly
  UserData
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import UserData from '../components/UserData.vue';

export default defineComponent({
  name: 'PageIndex',
  components: { UserData, QPage },
});
</script>
