<template lang="pug">
QPage.row.items-center.justify-evenly
  Login
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import Login from '../components/Login.vue';

export default defineComponent({
  name: 'PageIndex',
  components: { Login, QPage },
});
</script>
