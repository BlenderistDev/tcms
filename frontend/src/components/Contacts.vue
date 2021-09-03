<template lang="pug">
div {{ contacts }}
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';
import { ContactMap } from 'src/store/module-example/state';

export default defineComponent({
  name: 'Contacts',
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchContacts');
    const contacts = computed((): ContactMap|null => store.state.example.contacts);
    return {
      contacts,
    };
  },
});
</script>
