<template lang="pug">
div
  b-tabs
   b-tab(title="受注一覧" active)
    b-button(href="/order_info" variant="outline-info") じゅちゅう
   b-tab(title="仕入一覧")
    b-button(href="/purchase_info" variant="outline-info") しいれ
   b-tab(title="出荷一覧")
    b-button(href="/shipment_info" variant="outline-info") しゅっか
   b-tab(title="請求一覧")
    b-button(href="/invoice_info" variant="outline-info") せいきゅう
   b-tab(title="回収一覧")
        b-button(href="/recovery_info" variant="outline-info") かいしゅ
  logo-default
</template>

<script>
import LogoDefault from '~/components/Logo'
export default {
  fetch({ store, redirect }) {
    if (!store.state.auth) {
      return redirect('/login')
    }
  },
  components: {
    LogoDefault
  }
}
</script>
