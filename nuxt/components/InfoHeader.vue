<template lang="pug">
b-row.my-1(align-h="between")
  b-col(sm="3")
    b-button(variant="outline-info" block) 戻る
  b-col(sm="3")
    b-button(variant="outline-info" block) 戻る
  b-col(sm="3" v-if="info.next")
    b-button(:href="info.next.url" variant="outline-info" block) {{info.next.text}}へ
</template>

<script>
export default {
  props: ['info'],
}
</script>
