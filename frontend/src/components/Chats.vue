<template lang="pug">
QTable(
  title="Chats"
  :rows="rows"
  :columns="columns"
)
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';
import { QTable } from 'quasar';
import * as _ from 'lodash';
import { ChatMap, Chat } from './models';

export default defineComponent({
  name: 'Chats',
  components: {
    QTable,
  },
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchChats');
    const chats = computed((): ChatMap|null => store.state.example.chats);
    const rows = computed(() => _.values(chats.value));
    const columns = [
      {
        name: 'id',
        required: true,
        label: 'Id',
        align: 'left',
        field: (row: Chat) => row.id,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'creator',
        label: 'Creator',
        align: 'left',
        field: (row: Chat) => row.creator,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'date',
        label: 'Date',
        align: 'left',
        field: (row: Chat) => row.date,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'deactivated',
        label: 'Deactivated',
        align: 'left',
        field: (row: Chat) => row.deactivated,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'kicked',
        label: 'Kicked',
        align: 'left',
        field: (row: Chat) => row.kicked,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'left',
        label: 'Left',
        align: 'left',
        field: (row: Chat) => row.left,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'title',
        label: 'Title',
        align: 'left',
        field: (row: Chat) => row.title,
        format: (val: string) => `${val}`,
        sortable: true,
      },
    ];
    return {
      columns,
      rows,
    };
  },
});
</script>
