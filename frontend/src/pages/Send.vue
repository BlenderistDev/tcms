<template lang="pug">
QPage.row.items-center.justify-evenly
  SendForm
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { QPage } from 'quasar';
import SendForm from 'components/SendForm.vue';

export default defineComponent({
  name: 'PageIndex',
  components: { SendForm, QPage },
});
</script>
