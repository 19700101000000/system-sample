<template lang="pug">
div
  b-row.my-1
    b-col
      b-button(href="/order_info/new" variant="outline-primary" block) 新規見積
  b-row.my-1
    b-col
      b-card(no-body)
        b-tabs(card)
          b-tab(title="見積一覧" active)
            buy-list.my-1(:fields="fields" :items="items")
            b-button(href="/order_info" variant="outline-info") 見積

          b-tab(title="受注一覧")
            buy-list.my-1(:fields="fields" :items="items")
            b-button(href="/order_info" variant="outline-info") 受注

          b-tab(title="仕入一覧")
            buy-list.my-1(:fields="fields" :items="items")
            b-button(href="/purchase_info" variant="outline-info") 仕入

          b-tab(title="出荷一覧")
            buy-list.my-1(:fields="fields" :items="items")
            b-button(href="/shipment_info" variant="outline-info") 出荷

          b-tab(title="請求一覧")
            buy-list.my-1(:fields="fields" :items="items")
            b-button(href="/invoice_info" variant="outline-info") 請求

          b-tab(title="回収一覧")
            buy-list.my-1(:fields="fields" :items="items")
            b-button(href="/recovery_info" variant="outline-info") 回収
</template>

<style lang="sass" scope>
</style>

<script>
import BuyList from '~/components/List'
export default {
  fetch({ store, redirect }) {
    if (!store.state.auth) {
      return redirect('/login')
    }
  },
  components: {
    BuyList
  },
  data () {
    return {
      fields: [
        {
          key: 'id',
          label: '見積',
          sortable: true,
        },
        {
          key: 'client',
          label: '顧客',
          sortable: true,
        },
        {
          key: 'employee',
          label: '担当者',
          sortable: true,
        },
        {
          key: 'date',
          label: '作成日',
          sortable: true,
        },
      ],
      items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
      ]
    }
  }
}
</script>
