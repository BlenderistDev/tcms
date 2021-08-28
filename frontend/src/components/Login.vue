<template>
  <q-form>
    <q-input
      v-model="phone"
      type="tel"
    >
    </q-input>
  </q-form>
</template>

<script lang="ts">
import {
  defineComponent,
  ref,
} from 'vue';

export default defineComponent({
  name: 'CompositionComponent',
  setup() {
    const phone = ref('');
    return {
      phone,
    };
  },
});
</script>
