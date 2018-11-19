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
              :info="estimate_info"
              :fields="estimate_fields"
              :items="estimate_items"
            )

          b-tab(title="受注一覧")
            buy-list.my-1(
              :info="order_info"
              :fields="order_fields"
              :items="order_items"
            )

          b-tab(title="仕入一覧")
            buy-list.my-1(
              :info="purchase_info"
              :fields="purchase_fields"
              :items="purchase_items"
            )

          b-tab(title="出荷一覧")
            buy-list.my-1(
              :info="shipment_info"
              :fields="shipment_fields"
              :items="shipment_items"
            )

          b-tab(title="請求一覧")
            buy-list.my-1(
              :info="invoice_info"
              :fields="invoice_fields"
              :items="invoice_items"
            )

          b-tab(title="回収一覧")
            buy-list.my-1(
              :info="recovery_info"
              :fields="recovery_fields"
              :items="recovery_items"
              @click="getItems(5)"
            )
</template>

<style lang="sass" scope>
</style>

<script>
import axios from 'axios'
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
      estimate_info: {
        base_url: '/order_info/',
        radios: [
          {text: '見積中', value: 0},
          {text: '受注済', value: 1},
        ],
      },
      order_info: {
        base_url: '/order_info/',
        radios: [
          {text: '受注中', value: 0},
          {text: '仕入済', value: 1},
        ],
      },
      purchase_info: {
        base_url: '/purchase_info/',
        radios: [
          {text: '仕入中', value: 0},
          {text: '出荷済', value: 1},
        ],
      },
      shipment_info: {
        base_url: '/shipment_info/',
        radios: [
          {text: '出荷中', value: 0},
          {text: '請求済', value: 1},
        ],
      },
      invoice_info: {
        base_url: '/invoice_info/',
        radios: [
          {text: '請求中', value: 0},
          {text: '回収済', value: 1},
        ],
      },
      recovery_info: {
        base_url: '/recovery_info/',
        radios: [],
      },
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
      estimate_items: [],
      order_items: [],
      purchase_items: [],
      shipment_items: [],
      invoice_items: [],
      recovery_items: [],
    }
  },
  methods: {
    getItems: async function(index) {
      try {
        const { data } = await axios.get(`/api/display/${this.index}`)
        this.estimate_items = data
      } catch(error) {
        console.log(error.message)
      }
    }
  },
  mounted() {
    this.getItems()
  }
}
</script>
