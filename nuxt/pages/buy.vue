<template lang="pug">
div
  b-tabs
   b-tab(title="受注一覧" active)
    buy-list(:items="items")
    b-button(href="/order_info" variant="outline-info") 受注

   b-tab(title="仕入一覧")
    buy-list(:items="items")
    b-button(href="/purchase_info" variant="outline-info") 仕入

   b-tab(title="出荷一覧")
    buy-list(:items="items")
    b-button(href="/shipment_info" variant="outline-info") 出荷

   b-tab(title="請求一覧")
    buy-list(:items="items")
    b-button(href="/invoice_info" variant="outline-info") 請求

   b-tab(title="回収一覧")
    buy-list(:items="items")
    b-button(href="/recovery_info" variant="outline-info") 回収
</template>

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
      items: [
        {name: 'aaa', age: 10},
        {name: 'bbb', age: 20},
        {name: 'ccc', age: 30},
        {name: 'ddd', age: 40},
        {name: 'eee', age: 50},
        {name: 'fff', age: 60},
        {name: 'ggg', age: 10},
        {name: 'hhh', age: 20},
        {name: 'iii', age: 30},
        {name: 'jjj', age: 40},
        {name: 'kkk', age: 50},
        {name: 'lll', age: 60},
        {name: 'mmm', age: 10},
        {name: 'nnn', age: 20},
        {name: 'ooo', age: 30},
        {name: 'ppp', age: 40},
        {name: 'qqq', age: 50},
        {name: 'rrr', age: 60},
        {name: 'sss', age: 10},
        {name: 'ttt', age: 20},
        {name: 'uuu', age: 30},
        {name: 'vvv', age: 40},
        {name: 'www', age: 50},
        {name: 'xxx', age: 60},
      ]
    }
  }
}
</script>
