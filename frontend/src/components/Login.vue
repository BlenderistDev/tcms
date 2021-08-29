<template lang="pug">
QForm(@submit="login")
  QInput(
    v-model="phone"
    type="tel"
  )
  QBtn(type="submit") Submit
</template>

<script lang="ts">
import { QInput, QForm, QBtn } from 'quasar';
import {
  defineComponent,
  ref,
} from 'vue';
import { api } from 'src/boot/axios';
import { AxiosResponse } from 'axios';

export default defineComponent({
  name: 'CompositionComponent',
  components: {
    QInput,
    QForm,
    QBtn,
  },
  setup() {
    const phone = ref('');
    function login():Promise<AxiosResponse> {
      return api.post('/login', { phone: phone.value });
    }
    return {
      phone,
      login,
    };
  },
});
</script>
