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
            buy-list.my-1(
              base-url="/order_info/"
              :fields="estimate_fields"
              :items="estimate_items"
            )

          b-tab(title="受注一覧")
            buy-list.my-1(
              base-url="/order_info/"
              :fields="order_fields"
              :items="order_items"
            )

          b-tab(title="仕入一覧")
            buy-list.my-1(
              base-url="/purchase_info/"
              :fields="purchase_fields"
              :items="purchase_items"
            )

          b-tab(title="出荷一覧")
            buy-list.my-1(
              base-url="/shipment_info/"
              :fields="shipment_fields"
              :items="shipment_items"
            )

          b-tab(title="請求一覧")
            buy-list.my-1(
              base-url="/invoice_info/"
              :fields="invoice_fields"
              :items="invoice_items"
            )

          b-tab(title="回収一覧")
            buy-list.my-1(
              base-url="/recovery_info/"
              :fields="recovery_fields"
              :items="recovery_items"
            )
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
      estimate_fields: [
        {key: 'id', label: '見積', sortable: true},
        {key: 'client', label: '顧客', sortable: true},
        {key: 'employee', label: '担当者', sortable: true},
        {key: 'date', label: '作成日', sortable: true},
      ],
      order_fields: [
        {key: 'id', label: '受注', sortable: true},
        {key: 'client', label: '顧客', sortable: true},
        {key: 'employee', label: '担当者', sortable: true},
        {key: 'date', label: '作成日', sortable: true},
      ],
      purchase_fields: [
        {key: 'id', label: '仕入', sortable: true},
        {key: 'client', label: '顧客', sortable: true},
        {key: 'employee', label: '担当者', sortable: true},
        {key: 'date', label: '作成日', sortable: true},
      ],
      shipment_fields: [
        {key: 'id', label: '出荷', sortable: true},
        {key: 'client', label: '顧客', sortable: true},
        {key: 'employee', label: '担当者', sortable: true},
        {key: 'date', label: '作成日', sortable: true},
      ],
      invoice_fields: [
        {key: 'id', label: '請求', sortable: true},
        {key: 'client', label: '顧客', sortable: true},
        {key: 'employee', label: '担当者', sortable: true},
        {key: 'date', label: '作成日', sortable: true},
      ],
      recovery_fields: [
        {key: 'id', label: '回収', sortable: true},
        {key: 'client', label: '顧客', sortable: true},
        {key: 'employee', label: '担当者', sortable: true},
        {key: 'date', label: '作成日', sortable: true},
      ],
      estimate_items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '888888888', employee: { id: '222222222', name: 'bbb' }, client: { id: '444444444', name:  'BBB' }, date: '2018-11-22'},
      ],
      order_items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '888888888', employee: { id: '222222222', name: 'bbb' }, client: { id: '444444444', name:  'BBB' }, date: '2018-11-22'},
      ],
      purchase_items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '888888888', employee: { id: '222222222', name: 'bbb' }, client: { id: '444444444', name:  'BBB' }, date: '2018-11-22'},
      ],
      shipment_items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '888888888', employee: { id: '222222222', name: 'bbb' }, client: { id: '444444444', name:  'BBB' }, date: '2018-11-22'},
      ],
      invoice_items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '888888888', employee: { id: '222222222', name: 'bbb' }, client: { id: '444444444', name:  'BBB' }, date: '2018-11-22'},
      ],
      recovery_items: [
        {id: '999999999', employee: { id: '111111111', name: 'aaa' }, client: { id: '333333333', name:  'AAA' }, date: '2018-11-11'},
        {id: '888888888', employee: { id: '222222222', name: 'bbb' }, client: { id: '444444444', name:  'BBB' }, date: '2018-11-22'},
      ],
    }
  }
}
</script>
