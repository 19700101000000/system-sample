<template lang="pug">
  b-container
    b-card.my-1.mx-auto(
      bg-variant="light"
      title="Sign in"
      style="max-width: 30rem")
      b-form-group(
        horizontal
        label="Name:"
        label-class="text-sm-right"
        label-for="signin-name")
        b-form-input#signin-name(
          v-model="inName"
          type="text"
          placeholder="YOUR NAME")

      b-form-group(
        horizontal
        label="Password:"
        label-class="text-sm-right"
        label-for="signin-pass")
        b-form-input#signin-pass(
          v-model="inPass"
          type="password"
          placeholder="YOUR PASSWORD")

      p.text-center
        b-button(
          href="#"
          variant="outline-primary") Sign in
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"

@Component
export default class extends Vue {
  public inName: string = "";
  public inPass: string = "";
}
</script>
