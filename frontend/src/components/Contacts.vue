<template lang="pug">
QTable(
  title="Contacts"
  :rows="rows"
  :columns="columns"
)
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'src/store';
import { QTable } from 'quasar';
import * as _ from 'lodash';
import { ContactMap, Contact } from './models';

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
        field: (row: Contact) => row.id,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'Username',
        label: 'Username',
        align: 'left',
        field: (row: Contact) => row.username,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'Phone',
        label: 'Phone',
        align: 'left',
        field: (row: Contact) => row.phone,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'firstName',
        label: 'First name',
        align: 'left',
        field: (row: Contact) => row.firstName,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'lastName',
        label: 'Last name',
        align: 'left',
        field: (row: Contact) => row.lastName,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'mutualContact',
        label: 'Mutual contact',
        align: 'left',
        field: (row: Contact) => row.mutualContact,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'deleted',
        label: 'deleted',
        align: 'left',
        field: (row: Contact) => row.deleted,
        format: (val: string) => `${val}`,
        sortable: true,
      },
      {
        name: 'bot',
        label: 'bot',
        align: 'left',
        field: (row: Contact) => row.bot,
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
