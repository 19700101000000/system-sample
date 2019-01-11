<template lang="pug">
  b-container
    form-image
    card-image(v-for="value in values" :value="value")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import FormImage from "~/components/FormImage.vue"
import CardImage from "~/components/CardImage.vue"
import axios from "axios"

@Component({
  components: {
    FormImage,
    CardImage
  }
})
export default class extends Vue {
  public values = []

  public mounted() {
    axios.get("/api/get/galleries").then(({ data }) => {
      this.values = data.galleries;
    });
  }
}
</script>

<style lang="sass" scoped></style>
