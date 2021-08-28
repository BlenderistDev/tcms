<template lang="pug">
QForm
  QInput(
    v-model="phone"
    type="tel"
  )
</template>

<script lang="ts">
import { QInput, QForm } from 'quasar';
import {
  defineComponent,
  ref,
} from 'vue';

export default defineComponent({
  name: 'CompositionComponent',
  components: {
    QInput,
    QForm,
  },
  setup() {
    const phone = ref('');
    return {
      phone,
    };
  },
});
</script>
