<template lang="pug">
div
  b-row.my-1
    b-col(sm="4")
      b-form-input(type="number" v-model="num")
    b-col(sm="4")
      b-button(@click="getItems" variant="outline-success" block) 変更
    b-col(sm="4")
      b-button(@click="sample" variant="outline-info" block) サンプル
  b-row.my-1
    b-col
      b-table(striped outline :items="items")
</template>

<script>
import axios from 'axios'

export default {
  data() {
    return {
      num: 1,
      items: [],
    }
  },
  methods: {
    getItems: async function() {
      try {
        const { data } = await axios.get(`/api/sample/${this.num}`)
        this.items = data
      } catch(error) {
        console.log(error.message)
      }
    },
    sample() {
      alert('sample this')
    }
  },
  mounted() {
    this.getItems()
  }
}
</script>

