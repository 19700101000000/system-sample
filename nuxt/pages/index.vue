<template lang="pug">
  section
    b-container
      b-row.row
        b-col
          b-form(@submit="onSubmit")
            b-input-group(size="lg")
              b-form-input(
                size="sm",
                class="mr-sm-2",
                type="text",
                placeholder="キーワードを入力してください。",
                @change="handleChange"
                :value="query"
              )
              b-button(size="sm", class="my-2 my-sm-0", type="submit") 検索する
        b-col
      div.cards
        book-list-item(
          v-for='book in books',
          :key='book.id',
          :book='book',
        )
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import { State } from "vuex-class"
import BookListItem from "~/components/BookListItem.vue"

@Component({
  components: {
    BookListItem
  }
})
export default class extends Vue {
  @State books
  @State query

  handleChange(e): void {
    this.$store.commit('setQuery', e);
  }

  onSubmit(e): void {
    e.preventDefault();
    this.$store.dispatch('search');
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
