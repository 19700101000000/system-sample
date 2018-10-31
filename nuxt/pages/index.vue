<template lang="pug">
  b-container
    b-row
      b-col
        carousel
    b-row
      b-col
        p foo
        carousel
      b-col
        carousel
        p bar
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import Carousel from "~/components/Carousel.vue"

@Component({
  components: {
    Carousel
  }
})
export default class extends Vue {
}
</script>

<style lang="sass" scoped>
.row
  position: relative
.col
  margin: 1em 0
</style>
