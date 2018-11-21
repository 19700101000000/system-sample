<template lang="pug">
div
  input(v-model="message" plogreholder="edit")
  p message is:{{message}}

  input(type="checkbox" id="checkbox" v-model="checked")
  label(for="checkbox") {{checked}}
  input(type="text" v-if="checked")
  br
  template(v-if="toggle")
    label(for="username") Username
    input#username(type="text")
  template(v-else)
    label(for="email") Email
    input#email(type="text")
  br
  button(@click="toggle = !toggle") toggle
</template>


<script>
export default{
  data(){
    return{
      toggle: true,
      message: '',
      checked: false,
      username: '',
      email:'',
    }
  }
}
</script>
