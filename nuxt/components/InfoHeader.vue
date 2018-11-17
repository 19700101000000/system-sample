<template lang="pug">
b-row.my-1
  b-col(sm="3")
    b-button(variant="outline-info" block) 戻る
  b-col(sm="6")
    b-button(variant="outline-danger" block) 買注文
  b-col(sm="3" v-if="info.next")
    b-button(:href="info.next.url" variant="outline-success" block) {{info.next.text}}へ
</template>

<script>
export default {
  props: ['info'],
}
</script>
