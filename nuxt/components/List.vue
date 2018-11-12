<template lang="pug">
div
  b-table(striped hover :items="items.slice(currentPage * 10 - 10, currentPage * 10)")
  b-pagination-nav(align="center" base-url="#" :number-of-pages="parseInt(items.length / 10 + 1)" v-model="currentPage")
</template>
<script>
export default {
  props: ['items'],
  data () {
    return {
      currentPage: 1
    }
  }
}
</script>

<style lang="sass">
</style>
