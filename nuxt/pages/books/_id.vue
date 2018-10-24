
<template lang="pug">
  section
    div.cards
      book-detail-item(:book="book")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import BookDetailItem from "~/components/BookDetailItem.vue"
import { State } from "vuex-class"

@Component({
  components: {
    BookDetailItem
  }
})
export default class BookDetail extends Vue {
  @State book
  mounted() {
    this.$store.dispatch('fetchBook', this.$route.params.id)
  }
}
</script>
