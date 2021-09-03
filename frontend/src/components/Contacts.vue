<template lang="pug">
div {{ contacts }}
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';

export default defineComponent({
  name: 'Contacts',
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchContacts');
    const contacts = computed(() => store.state.example.contacts);
    return {
      contacts,
    };
  },
});
</script>
