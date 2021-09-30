<template lang="pug">
QForm(@submit="formSign" v-if="isSign")
  QInput(
    v-model="code"
    type="number"
    label="Code"
  )
  QBtn(type="submit") Submit
QForm(@submit="formLogin" v-else)
  QInput(
    v-model="phone"
    type="number"
    label="Phone"
  )
  QBtn(type="submit") Submit
</template>

<script lang="ts">
import { QInput, QForm, QBtn } from 'quasar';
import {
  defineComponent,
  ref
} from 'vue';
import { login, sign } from 'src/services/api';
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
    function formLogin():Promise<AxiosResponse> {
      isSign.value = true;
      return login(phone.value);
    }
    function formSign():Promise<AxiosResponse> {
      return sign(code.value);
    }
    return {
      phone,
      code,
      isSign,
      formLogin,
      formSign,
    };
  },
});
</script>
