<template lang="pug">
QForm(@submit="sign" v-if="isSign")
  QInput(
    v-model="code"
    type="number"
  )
  QBtn(type="submit") Submit
QForm(@submit="login" v-else)
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
    const code = ref('');
    const isSign = ref(false);
    function login():Promise<AxiosResponse> {
      isSign.value = true;
      return api.post('/login', { phone: phone.value });
    }
    function sign():Promise<AxiosResponse> {
      return api.post('/sign', { code: code.value });
    }
    return {
      phone,
      code,
      isSign,
      login,
      sign,
    };
  },
});
</script>
