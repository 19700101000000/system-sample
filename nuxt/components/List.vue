<template lang="pug">
div
  b-row.my-1
    b-col
      b-form-radio-group(
        v-model="radio_checked"
        buttons
        button-variant="outline-primary"
        :options="info.radios"
      )

  b-row.my-1
    b-col
      b-form-input(type="text" v-model="filter" placeholder="検索")
  b-row.my-1
    b-col
      b-table(
        striped
        hover
        fixed
        outlined
        :filter="filter"
        :fields="fields"
        :items="items"
        :per-page="per_page"
        :current-page="current_page"
        @filtered="onFiltered"
      )
        template(slot="id" slot-scope="data")
          a(:href="info.base_url + data.value") {{ data.value }}
        template(slot="client" slot-scope="data")
          a(:href="'#' + data.value.id") {{ data.value.name }}
        template(slot="employee" slot-scope="data")
          a(:href="'#' + data.value.id") {{ data.value.name }}
  b-row.my-1
    b-col
      b-pagination(
        align="center"
        :per-page="per_page"
        :total-rows="total_rows"
        v-model="current_page"
      )
</template>

<style lang="sass" scope>
</style>

<script>
export default {
  props: ['info', 'fields', 'items'],
  data () {
    return {
      radio_checked: 0,
      current_page: 1,
      filter: '',
      per_page: 10,
      total_rows: this.items.length
    }
  },
  methods: {
    onFiltered (filter_items) {
      this.total_rows = filter_items.length
      this.current_page = 1
    }
  }
}
</script>
