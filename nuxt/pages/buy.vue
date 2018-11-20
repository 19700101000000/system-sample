<template lang="pug">
div
  b-row.my-1
    b-col
      b-button(@click="sample" block) sample
    b-col
      b-button(href="/order_info/new" variant="outline-primary" block) 新規見積
  b-row.my-1
    b-col
      b-card(no-body)
        b-tabs(card)
          b-tab(title="見積一覧" @click="getItems(0)" active)
            buy-list.my-1(
              ref="checked_list"
              :info="estimate_info"
              :fields="estimate_fields"
              :items="estimate_items"
              @click="getItems(0)"
            )

          b-tab(title="受注一覧" @click="getItems(1)")
            buy-list.my-1(
              ref="checked_list"
              :info="order_info"
              :fields="order_fields"
              :items="order_items"
              @click="getItems(1)"
            )

          b-tab(title="仕入一覧" @click="getItems(2)")
            buy-list.my-1(
              ref="checked_list"
              :info="purchase_info"
              :fields="purchase_fields"
              :items="purchase_items"
              @click="getItems(2)"
            )

          b-tab(title="出荷一覧" @click="getItems(3)")
            buy-list.my-1(
              ref="checked_list"
              :info="shipment_info"
              :fields="shipment_fields"
              :items="shipment_items"
              @click="getItems(3)"
            )

          b-tab(title="請求一覧" @click="getItems(4)")
            buy-list.my-1(
              ref="checked_list"
              :info="invoice_info"
              :fields="invoice_fields"
              :items="invoice_items"
              @click="getItems(4)"
            )

          b-tab(title="回収一覧" @click="getItems(5)")
            buy-list.my-1(
              :info="recovery_info"
              :fields="recovery_fields"
              :items="recovery_items"
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

      estimate_id: 0,
      order_id: 1,
      purchase_id: 2,
      shipment_id: 3,
      invoice_id: 4,
      recovery_id: 5,
    }
  },
  methods: {
    getItems: async function(id) {
      try {
        let flag = 0
        flag = this.$refs.checked_list.$data.radio_checked
        const { data } = await axios.get(`/api/display/${id}?listflag=${flag}`)
        switch(id) {
        case this.estimate_id:
          this.estimate_items = data
          break;
        case this.order_id:
          this.order_items = data
          break;
        case this.purchase_id:
          this.purchase_items = data
          break;
        case this.shipment_id:
          this.shipment_items = data
          break;
        case this.invoice_id:
          this.invoice_items = data
          break;
        case this.recovery_id:
          this.recovery_items = data
          break;
        }
      } catch(error) {
        console.log(error.message)
      }
      console.log("foo")
    },
    sample() {
      alert(this.$refs.estimate_list.$data.radio_checked)
    }
  },
  mounted() {
    this.getItems(this.estimate_id)
    console.log("bar")
  }
}
</script>
