<template lang="pug">
b-row.my-1
  b-col(sm="3")
    b-button(@click="$router.go(-1)" variant="outline-info" block) 戻る
  b-col(sm="6")
    b-button(href="/buy" variant="outline-danger" block) 買注文
  b-col(sm="3" v-if="info.next")
    b-button(:href="info.next.url" variant="outline-success" block) {{info.next.text}}
</template>

<script>
export default {
  props: ['info'],
}
</script>
