<template lang="pug">
div {{ user }}
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';

export default defineComponent({
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
