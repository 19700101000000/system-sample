<template lang="pug">
  section
    b-container
      b-row.row
        b-col
          b-from
            b-input-group(size="lg")
              b-from-input(
                size="sm",
                class="mr-sm-2",
                type="text",
                placeholder="キーワードを入力してください。",
              )
              b-button(size="sm", class="my-2 my-sm-0", type="submit")
        b-col
      div.cards
        book-list-item
        book-list-item
        book-list-item
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import BookListItem from "~/components/BookListItem.vue"

@Component({
  components: {
    BookListItem
  }
})
export default class extends Vue {

}
</script>

<style lang="sass" scoped>
.cards
  display: flex
  flex-wrap: wrap

.row
  padding: 8px 0
</style>
