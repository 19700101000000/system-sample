<template lang="pug">
div
  b-row.my-1(align-h="center")
    b-col(sm="2")
      label(for="username") ID:
    b-col(sm="6")
      b-form-input#username(type="text" v-model="username" placeholder="YOUR ID")
  b-row.my-1(align-h="center")
    b-col(sm="2")
      label(for="password") PASS:
    b-col(sm="6")
      b-form-input#password(type="password" v-model="password" placeholder="YOUR PASSWORD")
  b-row.my-1(align-h="center")
    b-col(sm="4")
      b-button(variant="outline-danger" block) CANCEL
    b-col(sm="4")
      b-button(variant="outline-success" @click="login" block) LOGIN

</template>

<script>
export default {
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    async login() {
      await this.$store.dispatch('login', {
        username: this.username,
        password: this.password
      })
    }
  }
}
</script>

<style lang="sass" scope>
</style>
