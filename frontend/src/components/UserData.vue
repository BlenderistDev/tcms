<template lang="pug">
div {{ user }}
Contacts
</template>

<script lang="ts">
import Contacts from 'components/Contacts.vue';
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';

export default defineComponent({
  components: {
    Contacts,
  },
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchUser');
    const user = computed(() => store.state.example.user);
    return {
      user,
    };
  },
});
</script>
