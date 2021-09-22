<template lang="pug">
QForm(@submit="onSubmit")
  QSelect(
    v-model="formData.recipient"
    :options="contacts"
    label="Recipient"
    option-label="firstName"
    required
  )
  QInput(
    v-model="formData.message"
    label="Message"
    autogrow
    required
  )
  QBtn(type="submit") Send
</template>

<script lang="ts">
import { defineComponent, computed, reactive } from 'vue';
import { useStore } from 'src/store';
import {
  QForm,
  QSelect,
  QInput,
  QBtn
} from 'quasar';
import * as _ from 'lodash';
import { sendMessage, SendMessage } from 'src/services/api';
import { ContactMap } from './models';

export default defineComponent({
  name: 'SendForm',
  components: {
    QForm,
    QSelect,
    QInput,
    QBtn,
  },
  setup() {
    const store = useStore();
    void store.dispatch('example/fetchContacts');
    const contacts = computed((): ContactMap|null => _.values(store.state.example.contacts));

    const formData: SendMessage = reactive({
      recipient: null,
      message: '',
    });
    return {
      formData,
      contacts,
      onSubmit: () => {
        void sendMessage(formData);
      },
    };
  },
});
</script>
