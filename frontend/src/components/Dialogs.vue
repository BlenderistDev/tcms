<template lang="pug">
QTable(
  title="Dialogs"
  :rows="rows"
  :columns="columns"
)
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';
import { QTable } from 'quasar';
import * as _ from 'lodash';
import { DialogMap, Dialog } from './models';

export default defineComponent({
  name: 'Dialogs',
  components: {
    QTable,
  },
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchDialogs');
    const dialogs = computed((): DialogMap|null => store.state.example.dialogs);
    const rows = computed(() => _.values(dialogs.value));
    const columns = [
      {
        name: 'folderId',
        label: 'Folder id',
        align: 'left',
        field: (row: Dialog) => row.folderId,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'pinned',
        label: 'Pinned',
        align: 'left',
        field: (row: Dialog) => row.pinned,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'unreadCount',
        label: 'Unread count',
        align: 'left',
        field: (row: Dialog) => row.unreadCount,
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
