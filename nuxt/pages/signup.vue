<template lang="pug">
  b-container
    b-card.my-2.mx-auto(
      bg-variant="light"
      style="max-width: 30rem;"
      no-body)
      b-card-body
        h4 Sign up
        b-form-group(
          horizontal
          label="Name:"
          label-class="text-sm-right"
          label-for="inputName")
          b-form-input#inputName(
            v-model="name"
            :state="stateName"
            placeholder="Input Name.")
        b-form-group(
          horizontal
          label="Password:"
          label-class="text-sm-right"
          label-for="inputPassword")
          b-form-input#inputPassword(
            v-model="password"
            :state="statePassword"
            :disabled="disabled"
            type="password"
            placeholder="Input Password.")
        b-form-group(
          horizontal
          label="Retype:"
          label-class="text-sm-right"
          label-for="inputRetype")
          b-form-input#inputRetype(
            v-model="retype"
            :state="stateRetype"
            :disabled="disabled"
            type="password"
            placeholder="Retype Password.")
        p.text-danger(v-if="error") Fail Sign up.
        b-button(
          variant="outline-primary"
          :disabled="!stateName || !statePassword || !stateRetype || disabled"
          v-on:click="onSignup"
          block) Sign up
</template>

<script lang="ts">
import {
  Component,
  Vue
 } from "nuxt-property-decorator"
 import axios from "axios"

 @Component
export default class extends Vue {
  public name      = "";
  public password  = "";
  public retype    = "";

  public error = false;
  public disabled = false;

  public get stateName(): boolean {
    return this.name.length > 0;
  }
  public get statePassword(): boolean {
    return this.password.length > 0;
  }
  public get stateRetype(): boolean {
    return this.password === this.retype && this.retype.length > 0;
  }

  public onSignup() {
    this.disabled = true;
    axios.post("/api/auth/signup", {
      name: this.name,
      password: this.password,
    }).then(({ data }) => {
      this.error = false;
      window.document.location.href = "/";
    }).catch((result) => {
      this.error = true;
      this.disabled = false;
    });
  }
}
</script>
