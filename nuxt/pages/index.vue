<template lang="pug">
  b-container
    b-row
      b-col
        ad-carousel
    b-row
      b-col
        rank-table
      b-col
        rank-table
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import AdCarousel from "~/components/AdCarousel.vue"
import RankTable from "~/components/RankTable.vue"

@Component({
  components: {
    AdCarousel,
    RankTable
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
