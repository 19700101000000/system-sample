<template lang="pug">
div.card
  div.thumbnail
    b-img(:src="thumbnail", fluid, alt="Responsive image")
  div.content
    h2 {{ title }}
    p.card-text {{ description }}
</template>

<script lang="ts">
import {
  Component,
  Prop,
  Vue
} from "nuxt-property-decorator"

@Component({})
export default class BookDetailItem extends Vue {
  @Prop() book
  get title() {
    try {
      return this.book.volumeInfo.title;
    } catch {}
    return ''
  }
  get description() {
    try {
      return this.book.volumeInfo.description
    } catch {}
  }

  get thumbnail(): string | boolean {
    try {
      return this.book.volumeInfo.imageLinks.thumbnail
    } catch {}
    return ''
  }
}
</script>

<style lang="sass" scoped>
.cards
  display: flex
  flex-wrap: wrap

.row
  padding: 8px 0
</style>
