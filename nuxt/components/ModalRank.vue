<template lang="pug">
  b-modal(
    ref="modal"
    title="Rates"
    ok-only
    ok-variant="outline-primary"
    size="lg")
    b-card.mt-2(v-for="rate in rates" :title="`${rate.observer} (${rate.rate})`") {{ rate.review }}
</template>

<script lang="ts">
import {
  Vue, 
  Component,
  Prop,
} from "nuxt-property-decorator"
import axios from "axios"

interface rate {
  observer: String,
  rate:     number,
  review:   string,
}

@Component
export default class ModalRank extends Vue {
  
  @Prop() userName: String;

  public rates: rate[] = [];

  public showModal() {
    const modal: any  = this.$refs.modal;
    modal.show();

    axios.get("/api/get/rates/" + this.userName).then(({ data }) => {
      this.rates = data.rates;
    }).catch((result) => {
    });
  }
}
</script>
