<template lang="pug">
QTable(
  title="Treats"
  :rows="rows"
  :columns="columns"
)
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';
import { ContactMap, Contact } from 'src/store/module-example/state';
import { QTable } from 'quasar';
import * as _ from 'lodash';

export default defineComponent({
  name: 'Contacts',
  components: {
    QTable,
  },
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchContacts');
    const contacts = computed((): ContactMap|null => store.state.example.contacts);
    const rows = computed(() => _.values(contacts.value));
    const columns = [
      {
        name: 'Id',
        required: true,
        label: 'Id',
        align: 'left',
        field: (row: Contact) => row.Id,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'Username',
        label: 'Username',
        align: 'left',
        field: (row: Contact) => row.Username,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'Phone',
        label: 'Phone',
        align: 'left',
        field: (row: Contact) => row.Phone,
        format: (val: string) => `${val}`,
        sortable: true,
      },
    ];
    return {
      columns,
      contacts,
      rows,
    };
  },
});
</script>
